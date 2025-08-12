# Send Buttons

Send a message containing interactive reply buttons.

**Endpoint:** `/chat/send/buttons`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | Recipient JID or phone number. |
| `Title` | string | yes | Title displayed above the buttons. |
| `Buttons` | array[object] | yes | Up to 3 buttons with `ButtonId` and `ButtonText`. |
| `Id` | string | no | Custom message ID. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","Title":"Escolha","Buttons":[{"ButtonId":"1","ButtonText":"Sim"},{"ButtonId":"2","ButtonText":"Não"}]}' \
  http://localhost:8080/chat/send/buttons
```
