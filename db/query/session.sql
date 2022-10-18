-- name: CreateSession :one
INSERT INTO sessions (
username,
refresh_token_uuid,
refresh_token,
user_agent,
client_ip,
expires_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetActiveSession :one
SELECT * FROM sessions
WHERE
refresh_token_uuid = coalesce(sqlc.narg('refresh_token_uuid'), refresh_token_uuid)
AND refresh_token = coalesce(sqlc.narg('refresh_token'), refresh_token)
AND username = coalesce(sqlc.narg('username'), username)
AND client_ip = coalesce(sqlc.narg('client_ip'), client_ip)
AND user_agent = coalesce(sqlc.narg('user_agent'), user_agent)
AND status = true;