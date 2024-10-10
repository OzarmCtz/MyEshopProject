-- name: GetItemCategoryLink :one
SELECT * FROM `items_category_link`
WHERE icl_id = ? LIMIT 1;

-- name: GetItemCategoryLinkBySubCategory :one
SELECT 
    ic.ic_id, 
    ic.ic_name, 
    ic.ic_description, 
    ic.ic_picture_url
FROM 
    `items_category` ic
JOIN 
    `items_category_link` icl ON ic.ic_id = icl.icl_items_category_id
WHERE 
    icl.icl_items_sub_category_id = ? 
LIMIT 1;



-- name: ListItemsCategoryLinkByCategory :many
SELECT * FROM `items_category_link`
WHERE icl_items_category_id = ? ORDER BY icl_id ASC;


-- name: ListItemsCategoryLink :many
SELECT * FROM `items_category_link`
ORDER BY icl_id ASC;


-- name: CreateItemsCategoryLink :execresult
INSERT INTO `items_category_link` (icl_items_sub_category_id , icl_items_category_id
) VALUES (? , ?);

-- name: DeleteItemsCategoryLink :execrows
DELETE FROM `items_category_link`
WHERE icl_id = ?;

-- name: UpdateItemsCategoryLink :execrows
UPDATE `items_category_link` SET icl_items_sub_category_id = ? , icl_items_category_id = ?
WHERE icl_id = ?;



-- name: UpdateItemsCategoryLinkBySubCategory :execrows
UPDATE `items_category_link` SET  icl_items_category_id = ?
WHERE icl_items_sub_category_id = ?;