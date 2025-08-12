# Download Image

Download an image from WhatsApp servers using message metadata.

**Endpoint:** `/chat/downloadimage`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Url` | string | no | Download URL provided by WhatsApp. |
| `DirectPath` | string | no | Direct path to the file. |
| `MediaKey` | string | no | Base64-encoded media key. |
| `Mimetype` | string | no | MIME type of the file. |
| `FileEncSHA256` | string | no | SHA256 of the encrypted file. |
| `FileSHA256` | string | no | SHA256 of the decrypted file. |
| `FileLength` | number | no | File size in bytes. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Url":"https://mmg.whatsapp.net/d/f/ABCDEF","DirectPath":"/v/t62.7118-24/...","MediaKey":"base64==","Mimetype":"image/jpeg","FileEncSHA256":"base64==","FileSHA256":"base64==","FileLength":12345}' \
  http://localhost:8080/chat/downloadimage
```
