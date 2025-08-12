# Download Document

Download a document from WhatsApp servers using message metadata.

**Endpoint:** `/chat/downloaddocument`

**Method:** `POST`

The payload fields are the same as [Download Image](download-image.md).

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Url":"https://mmg.whatsapp.net/d/f/ABCDEF","DirectPath":"/v/t62.7118-24/...","MediaKey":"base64==","Mimetype":"application/pdf"}' \
  http://localhost:8080/chat/downloaddocument
```
