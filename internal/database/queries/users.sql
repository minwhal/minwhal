-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = ? limit 1;

-- name: SaveUser: one
INSERT INTO users (id, hashed_password, email)
VALUES(?,?,?)
RETURNING *;
