# Get User Info

Retrieve detailed information about one or more users.

**Endpoint:** `/user/info`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `Phone` | array[string] | yes | List of phone numbers or JIDs to query. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"Phone":["5491155553935"]}' \
  http://localhost:8080/user/info
```
