-- name: GetDiscount :one
SELECT * FROM discount
WHERE d_id = ? LIMIT 1;


-- name: GetDiscountByCode :one
SELECT * FROM discount
WHERE d_code = ? LIMIT 1;

-- name: ListDiscounts :many
SELECT * FROM discount
ORDER BY d_id ASC;

-- name: CreateDiscount :execresult
INSERT INTO discount (d_code ,d_description, d_start_time  , d_end_time , d_zone_time , d_is_disabled , d_price_type , d_value ) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteDiscount :execrows
DELETE FROM discount
WHERE d_id = ?;

-- name: UpdateDiscount :execrows
UPDATE discount SET d_code = ?, d_description = ?, d_start_time = ?, d_end_time = ?, d_zone_time = ?, d_is_disabled = ?, d_price_type = ?, d_value = ?
WHERE d_id = ?;