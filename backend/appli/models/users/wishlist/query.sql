-- name: GetUserWishList :one

SELECT * FROM users_wishlist
WHERE wl_id = ? LIMIT 1;


-- name: ListUserWishListByUser :many

SELECT * FROM users_wishlist
WHERE wl_user_id = ?;

-- name: CreateUserWishList :execresult
INSERT INTO users_wishlist (wl_user_id, wl_items_id ,wl_times_added 
) VALUES (?, ?, ?);


-- name: DeleteUserWishListByUserAndItems :execrows
DELETE FROM users_wishlist
WHERE wl_user_id = ? AND wl_items_id = ?;