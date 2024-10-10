-- name: GetItemReview :one
SELECT * FROM items_reviews
WHERE ir_id = ? LIMIT 1;


-- name: ListItemsReviews :many
SELECT * FROM items_reviews
ORDER BY ir_id ASC;


-- name: ListItemsReviewsByItemId :many
SELECT * FROM items_reviews
WHERE ir_items_id = ? ;


-- name: ListItemsReviewsByUserId :many
SELECT * FROM items_reviews
WHERE ir_user_id = ?;

-- name: CreateItemReview :execresult
INSERT INTO items_reviews (ir_user_id, ir_items_id ,ir_comments, ir_stars
) VALUES (?, ?, ?, ?);

-- name: DeleteItemReview :execrows
DELETE FROM items_reviews
WHERE ir_id = ?;

