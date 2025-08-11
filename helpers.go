package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"go.mau.fi/whatsmeow"
	"net/http"
	"os"
	"regexp"
)

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Update entry in User map
func updateUserInfo(values interface{}, field string, value string) interface{} {
	log.Debug().Str("field", field).Str("value", value).Msg("User info updated")
	values.(Values).m[field] = value
	return values
}

// webhook for regular messages
func callHook(myurl string, payload map[string]string, id string) {
	log.Info().Str("url", myurl).Msg("Sending POST to client " + id)

	// Log the payload map
	log.Debug().Msg("Payload:")
	for key, value := range payload {
		log.Debug().Str(key, value).Msg("")
	}

	client := clientManager.GetHTTPClient(id)

	format := os.Getenv("WEBHOOK_FORMAT")
	if format == "json" {
		// Send as pure JSON
		// The original payload is a map[string]string, but we want to send the postmap (map[string]interface{})
		// So we try to decode the jsonData field if it exists, otherwise we send the original payload
		var body interface{} = payload
		if jsonStr, ok := payload["jsonData"]; ok {
			var postmap map[string]interface{}
			err := json.Unmarshal([]byte(jsonStr), &postmap)
			if err == nil {
				postmap["token"] = payload["token"]
				body = postmap
			}
		}
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(myurl)
		if err != nil {
			log.Debug().Str("error", err.Error())
		}
	} else {
		// Default: send as form-urlencoded
		_, err := client.R().SetFormData(payload).Post(myurl)
		if err != nil {
			log.Debug().Str("error", err.Error())
		}
	}
}

// webhook for messages with file attachments
func callHookFile(myurl string, payload map[string]string, id string, file string) error {
	log.Info().Str("file", file).Str("url", myurl).Msg("Sending POST")

	client := clientManager.GetHTTPClient(id)

	// Create final payload map
	finalPayload := make(map[string]string)
	for k, v := range payload {
		finalPayload[k] = v
	}

	finalPayload["file"] = file

	log.Debug().Interface("finalPayload", finalPayload).Msg("Final payload to be sent")

	resp, err := client.R().
		SetFiles(map[string]string{
			"file": file,
		}).
		SetFormData(finalPayload).
		Post(myurl)

	if err != nil {
		log.Error().Err(err).Str("url", myurl).Msg("Failed to send POST request")
		return fmt.Errorf("failed to send POST request: %w", err)
	}

	log.Debug().Interface("payload", finalPayload).Msg("Payload sent to webhook")
	log.Info().Int("status", resp.StatusCode()).Str("body", string(resp.Body())).Msg("POST request completed")

	return nil
}

func (s *server) respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Error().Err(err).Msg("Failed to encode JSON response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

var mentionRegex = regexp.MustCompile(`@([0-9]+)`)

// replaceAtMentions scans text for @phone patterns, gathers the numbers found
// and returns the original text so @ references remain clickable in clients.
func replaceAtMentions(text string, _ *whatsmeow.Client) (string, []string) {
	matches := mentionRegex.FindAllStringSubmatchIndex(text, -1)
	if matches == nil {
		return text, nil
	}

	phones := make([]string, 0, len(matches))
	for _, m := range matches {
		phone := text[m[2]:m[3]]
		if _, ok := parseJID(phone); !ok {
			continue
		}
		phones = append(phones, phone)
	}

	return text, phones
}

// ProcessOutgoingMedia handles media processing for outgoing messages with S3 support
func ProcessOutgoingMedia(userID string, contactJID string, messageID string, data []byte, mimeType string, fileName string, db *sqlx.DB) (map[string]interface{}, error) {
	// Check if S3 is enabled for this user
	var s3Config struct {
		Enabled       bool   `db:"s3_enabled"`
		MediaDelivery string `db:"media_delivery"`
	}
	err := db.Get(&s3Config, "SELECT s3_enabled, media_delivery FROM users WHERE id = $1", userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get S3 config")
		s3Config.Enabled = false
		s3Config.MediaDelivery = "base64"
	}

	// Process S3 upload if enabled
	if s3Config.Enabled && (s3Config.MediaDelivery == "s3" || s3Config.MediaDelivery == "both") {
		// Process S3 upload (outgoing messages are always in outbox)
		s3Data, err := GetS3Manager().ProcessMediaForS3(
			context.Background(),
			userID,
			contactJID,
			messageID,
			data,
			mimeType,
			fileName,
			false, // isIncoming = false for sent messages
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload media to S3")
			// Continue even if S3 upload fails
		} else {
			return s3Data, nil
		}
	}

	return nil, nil
}

// addStickerMetadata embeds a fixed author and pack name into a WebP sticker.
// It injects a minimal EXIF block so WhatsApp displays the pack information.
func addStickerMetadata(data []byte, packName, author string) []byte {
	json := fmt.Sprintf("{\"sticker-pack-id\":\"%s\",\"sticker-pack-name\":\"%s\",\"sticker-pack-publisher\":\"%s\"}", "com.botadmin", packName, author)
	jsonb := append([]byte(json), 0x00)

	exif := []byte{
		0x49, 0x49, 0x2A, 0x00,
		0x08, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x41, 0x57,
		0x07, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	binary.LittleEndian.PutUint32(exif[14:], uint32(len(jsonb)))
	exif = append([]byte("Exif\x00\x00"), append(exif, jsonb...)...)

	size := make([]byte, 4)
	binary.LittleEndian.PutUint32(size, uint32(len(exif)))
	chunk := append(append([]byte("EXIF"), size...), exif...)
	if len(chunk)%2 == 1 {
		chunk = append(chunk, 0)
	}

	out := append(data[:12], append(chunk, data[12:]...)...)
	binary.LittleEndian.PutUint32(out[4:8], uint32(len(out)-8))
	return out
}
