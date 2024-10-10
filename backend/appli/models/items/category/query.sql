-- name: GetItemCategory :one
SELECT * FROM items_category
WHERE ic_id = ? LIMIT 1;

-- name: ListItemsCategory :many
SELECT * FROM items_category
ORDER BY ic_id ASC;


-- name: ListItemsCategoryAndOccurence :many
SELECT 
    ic.ic_id,
    ic.ic_name,
    ic.ic_description,
    ic.ic_picture_url,
    COUNT(DISTINCT isc.isc_id) AS ic_on_isc,
    COUNT(DISTINCT i.i_id) AS total_items_count
FROM 
    items_category ic
LEFT JOIN 
    items_category_link icl ON ic.ic_id = icl.icl_items_category_id
LEFT JOIN 
    items_sub_category isc ON icl.icl_items_sub_category_id = isc.isc_id
LEFT JOIN 
    items_sub_category_link iscl ON isc.isc_id = iscl.iscl_sub_category_id
LEFT JOIN 
    items i ON iscl.iscl_items_id = i.i_id
GROUP BY 
    ic.ic_id, ic.ic_name, ic.ic_description, ic.ic_picture_url
ORDER BY 
    ic.ic_id ASC;




-- name: CreateItemCategory :execresult
INSERT INTO items_category (ic_name, ic_description , ic_picture_url
) VALUES (?, ? ,?);

-- name: DeleteItemCategory :execrows
DELETE FROM items_category
WHERE ic_id = ?;

-- name: UpdateItemCategory :execrows
UPDATE items_category SET ic_name = ?, ic_description = ? , ic_picture_url = ? WHERE ic_id = ?;