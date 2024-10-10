-- name: GetItem :one
SELECT * FROM items
WHERE i_id = ? LIMIT 1;

-- name: ListActiveItems :many
SELECT * FROM items
WHERE i_is_disabled = ? ;


-- name: ListItems :many
SELECT 
    i.*,
    isc.isc_name AS sub_category_name
FROM 
    items i
LEFT JOIN 
    items_sub_category_link iscl ON i.i_id = iscl.iscl_items_id
LEFT JOIN 
    items_sub_category isc ON iscl.iscl_sub_category_id = isc.isc_id
ORDER BY 
    i.i_id ASC;


-- name: CreateItem :execresult
INSERT INTO items (i_title, i_description ,i_price, i_quantity ,  i_picture_url , i_file_path , i_is_disabled , i_release_date
) VALUES (?, ? , ?, ?, ?, ?, ?, ?);

-- name: DeleteItem :execrows
DELETE FROM items
WHERE i_id = ?;

-- name: UpdateItem :execrows
UPDATE items SET i_title = ?,i_description = ?, i_price = ?, i_quantity = ?  , i_picture_url = ? , i_file_path = ? , i_is_disabled =  ? , i_release_date = ? WHERE i_id = ?;