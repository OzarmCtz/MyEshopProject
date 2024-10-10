-- name: GetItemsSubCategoryLink :one
SELECT * FROM `items_sub_category_link`
WHERE iscl_id = ? LIMIT 1;


-- name: GetItemsSubCategoryLinkByItem :one
SELECT 
    isc.isc_id,
    isc.isc_name,
    isc.isc_description,
    isc.isc_picture_url
FROM 
    items_sub_category_link iscl
JOIN 
    items_sub_category isc ON iscl.iscl_sub_category_id = isc.isc_id
WHERE 
    iscl.iscl_items_id = ? 
LIMIT 1;


-- name: ListItemsSubCategoryLinkByCategory :many
SELECT * FROM `items_sub_category_link`
WHERE iscl_sub_category_id = ? ORDER BY iscl_id ASC;


-- name: ListItemsSubCategoryLink :many
SELECT * FROM `items_sub_category_link`
ORDER BY iscl_id ASC;


-- name: CreateItemsSubCategoryLink :execresult
INSERT INTO `items_sub_category_link` (iscl_items_id , iscl_sub_category_id
) VALUES (? , ?);

-- name: CreateItemsSubCategoryLinkBySubCategoryName :execresult
INSERT INTO `items_sub_category_link` (iscl_items_id, iscl_sub_category_id)
SELECT ?, isc.isc_id FROM `items_sub_category` isc
WHERE isc.isc_name = ?;

-- name: DeleteItemsSubCategoryLink :execrows
DELETE FROM `items_sub_category_link`
WHERE iscl_id = ?;

-- name: UpdateItemsSubCategoryLink :execrows
UPDATE `items_sub_category_link` SET iscl_items_id = ? , iscl_sub_category_id = ?
WHERE iscl_id = ?;


-- name: UpdateItemsSubCategoryLinkByItem :execrows
UPDATE `items_sub_category_link` SET iscl_sub_category_id = ?
WHERE iscl_items_id = ?;