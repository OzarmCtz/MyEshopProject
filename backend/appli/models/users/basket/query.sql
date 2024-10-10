-- name: GetUserBasket :one

SELECT * FROM users_basket
WHERE ub_id = ? LIMIT 1;

-- name: ListUserBasketByUser :many

SELECT * FROM users_basket
WHERE ub_user_id = ?;

-- name: CreateUserBasket :execresult
INSERT INTO users_basket (ub_user_id, ub_items_id ,ub_time_added , ub_quantity
) VALUES (?, ?, ? , ?);


-- name: DeleteUserBasketByUserAndItems :execrows
DELETE FROM users_basket
WHERE ub_user_id = ? AND ub_items_id = ?;