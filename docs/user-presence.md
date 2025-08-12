# Set User Presence

Set the global online/offline presence for the connected account.

**Endpoint:** `/user/presence`

**Method:** `POST`

## Payload

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | string | yes | `available` to appear online or `unavailable` to appear offline. |

## Example

```bash
curl -X POST -H 'Token: 1234ABCD' -H 'Content-Type: application/json' \
  --data '{"type":"available"}' \
  http://localhost:8080/user/presence
```
