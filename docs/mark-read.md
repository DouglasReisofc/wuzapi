# Mark Messages as Read

Marks one or more messages as read in a chat or group.

**Endpoint:** `/chat/markread`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Chat` | string | yes | JID of the chat or group containing the messages. |
| `Id` | array[string] | yes | Message IDs to mark as read. |
| `Sender` | string | required for groups | Author JID when marking messages in a group. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Chat":"120363025945853632@g.us","Id":["AABBCC112233"],"Sender":"5491155553935@s.whatsapp.net"}' \
  http://localhost:8080/chat/markread
```
