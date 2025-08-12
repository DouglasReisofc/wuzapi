# Chat Presence

Update the typing or recording state for a specific chat.

**Endpoint:** `/chat/presence`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | JID or phone number of the chat. |
| `State` | string | yes | Presence state such as `composing`, `recording`, or `paused`. |
| `Media` | string | no | Optional media type, e.g. `audio` or `video`. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","State":"composing"}' \
  http://localhost:8080/chat/presence
```
