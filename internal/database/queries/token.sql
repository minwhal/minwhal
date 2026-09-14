-- name: GetToken :one
SELECT * FROM refresh_tokens
WHERE token = ? limit 1;

-- name: DeleteTokenByUserID :exec
DELETE FROM refresh_tokens 
WHERE user_id = ?;

-- name: CreateToken :one
INSERT INTO refresh_tokens  (id, expires_at, token, user_id)
VALUES(?,?,?,?)
RETURNING *;