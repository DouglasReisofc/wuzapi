# React to Message

Send or remove an emoji reaction on a message.

**Endpoint:** `/chat/react`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | JID or phone number of the chat. |
| `Body` | string | yes | Emoji to react with. Use `"remove"` to delete the reaction. |
| `Id` | string | yes | Message ID to react to. Prefix with `me:` if the message is yours. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","Body":"👍","Id":"ABCD1234"}' \
  http://localhost:8080/chat/react
```
