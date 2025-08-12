# Send Poll

Create a poll in a group chat.

**Endpoint:** `/chat/send/poll`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Group` | string | yes | Group JID where the poll will be sent. |
| `Header` | string | yes | Poll question or header text. |
| `Options` | array[string] | yes | List of poll options (minimum of 2). |
| `Id` | string | no | Custom message ID. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Group":"120363025945853632@g.us","Header":"Votar","Options":["Sim","Não"]}' \
  http://localhost:8080/chat/send/poll
```
