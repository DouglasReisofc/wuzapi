# Send List

Send an interactive list message with optional sections.

**Endpoint:** `/chat/send/list`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | Recipient JID or phone number. |
| `ButtonText` | string | yes | Text of the button that opens the list. |
| `Desc` | string | yes | Description shown in the message. |
| `TopText` | string | yes | Header text at the top of the list. |
| `Sections` | array[object] | no | List sections with `Title` and an array of `Rows`. |
| `FooterText` | string | no | Optional footer text. |
| `Id` | string | no | Custom message ID. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","ButtonText":"Abrir lista","Desc":"Escolha uma opção","TopText":"Opções","Sections":[{"Title":"Seção 1","Rows":[{"Title":"Item 1","Desc":"Primeiro","RowId":"1"}]}]}' \
  http://localhost:8080/chat/send/list
```
