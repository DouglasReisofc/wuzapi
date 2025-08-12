# Get Avatar

Fetch the profile picture for a user or group.

**Endpoint:** `/user/avatar`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | string | yes | JID or phone number to fetch the avatar for. |
| `Preview` | bool | no | If `true`, returns the thumbnail image. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":"5491155553935","Preview":false}' \
  http://localhost:8080/user/avatar
```
