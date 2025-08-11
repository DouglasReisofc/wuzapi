# Delete Message

Delete a previously sent message from a chat or group.

**Endpoint:** `/chat/delete`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | JID of the chat or group. |
| `Id` | string | yes | Message ID to revoke. |
| `Participant` | string | required for groups | Phone or JID of the sender when deleting someone else's message in a group. |

When `Phone` refers to a group (`@g.us`), the `Participant` field must contain the sender's JID. This allows group admins to delete messages from any participant.

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"120363025945853632@g.us","Id":"AABBCC11223344","Participant":"5491155553935"}' \
  http://localhost:8080/chat/delete
```
