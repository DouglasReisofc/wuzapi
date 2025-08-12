# Edit Message

Replace the body of an already sent text message.

**Endpoint:** `/chat/send/edit`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | JID or phone number of the chat. |
| `Body` | string | yes | New message text. |
| `Id` | string | yes | ID of the message to edit. |
| `ContextInfo` | object | no | Context for replying or mentioning users. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","Body":"Texto editado","Id":"ABCD1234"}' \
  http://localhost:8080/chat/send/edit
```
