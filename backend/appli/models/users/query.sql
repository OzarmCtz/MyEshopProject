-- name: GetUser :one
SELECT * FROM users
WHERE u_id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE u_email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY u_id ASC;

-- name: CreateUser :execresult
INSERT INTO users (u_email, u_uid ,u_register_date, u_is_disabled 
) VALUES (?, ? , ?, ?);

-- name: DeleteUser :execrows
DELETE FROM users
WHERE u_id = ?;

-- name: UpdateUser :execrows
UPDATE users SET u_email = ?,u_uid = ?, u_register_date = ?, u_is_disabled = ? WHERE u_id = ?;