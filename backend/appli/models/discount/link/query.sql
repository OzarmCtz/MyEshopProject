-- name: GetDiscountLink :one
SELECT * FROM discount_link
WHERE dl_id = ? LIMIT 1;


-- name: ListDiscountsLinks :many
SELECT * FROM discount_link
ORDER BY dl_id ASC;

-- name: CreateDiscountLink :execresult
INSERT INTO discount_link (dl_discount_id ,dl_items_id, dl_items_sub_category  , dl_items_category) 
VALUES (?, ?, ?, ?);

-- name: DeleteDiscountLink :execrows
DELETE FROM discount_link
WHERE dl_id = ?;

-- name: UpdateDiscountLink :execrows
UPDATE discount_link
SET dl_discount_id = ?, dl_items_id = ?, dl_items_sub_category = ?, dl_items_category = ?
WHERE dl_id = ?;




