-- name: GetItemSubCategory :one
SELECT * FROM items_sub_category
WHERE isc_id = ? LIMIT 1;

-- name: GetItemSubCategoryAndCategoryLinked :one
SELECT 
    isc.isc_id,
    isc.isc_name,
    isc.isc_description,
    isc.isc_picture_url,
    ic.ic_name
FROM 
    items_sub_category isc
LEFT JOIN 
    items_category_link icl ON isc.isc_id = icl.icl_items_sub_category_id
LEFT JOIN 
    items_category ic ON icl.icl_items_category_id = ic.ic_id
WHERE 
    isc.isc_id = ? 
LIMIT 1;

-- name: ListItemSubCategoryAndCategoryLinked :many
SELECT 
    isc.isc_id,
    isc.isc_name,
    isc.isc_description,
    isc.isc_picture_url,
    ic.ic_name,
    COUNT(iscl.iscl_items_id) AS item_count
FROM 
    items_sub_category isc
LEFT JOIN 
    items_category_link icl ON isc.isc_id = icl.icl_items_sub_category_id
LEFT JOIN 
    items_category ic ON icl.icl_items_category_id = ic.ic_id
LEFT JOIN 
    items_sub_category_link iscl ON isc.isc_id = iscl.iscl_sub_category_id
GROUP BY 
    isc.isc_id,
    ic.ic_name
ORDER BY 
    isc.isc_id;


-- name: ListItemsSubCategory :many
SELECT * FROM items_sub_category
ORDER BY isc_id ASC;

-- name: CreateItemSubCategory :execresult
INSERT INTO items_sub_category (isc_name, 
isc_description , isc_picture_url) VALUES (?, ? , ?);

-- name: DeleteItemSubCategory :execrows
DELETE FROM items_sub_category
WHERE isc_id = ?;

-- name: UpdateItemSubCategory :execrows
UPDATE items_sub_category SET isc_name = ?, isc_description  = ? , isc_picture_url = ? WHERE isc_id = ?;
