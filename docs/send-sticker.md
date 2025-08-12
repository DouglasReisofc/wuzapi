# Send Sticker

Send a static or animated sticker to a contact or group.

**Endpoint:** `/chat/send/sticker`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | Recipient JID or phone number. |
| `Sticker` | string | yes | Base64 data URL or http(s) link to a WebP image or MP4/WebM clip. Use `data:image/webp;base64,` for static stickers and `data:video/mp4;base64,` (or `data:video/webm;base64,`) for animated ones. |
| `Mentions` | array[string] | no | List of numbers to mention. |
| `Id` | string | no | Custom message ID. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","Sticker":"data:video/mp4;base64,AAA..."}' \
  http://localhost:8080/chat/send/sticker
```
