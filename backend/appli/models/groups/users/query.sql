-- name: GetGroupUser :one
SELECT * FROM `groups_users`
WHERE gu_id = ? LIMIT 1;

-- name: GetGroupByName :one
SELECT * FROM `groups_users`
WHERE gu_name = ? LIMIT 1;

-- name: ListGroupsUsers :many
SELECT * FROM `groups_users`
ORDER BY gu_id ASC;


-- name: ListGroupsUserByUser :many
SELECT gu.* FROM `groups_users` gu
INNER JOIN `groups_users_link` gul ON gul.gul_group_id = gu.gu_id
WHERE gul.gul_user_id = ? ORDER BY gu.gu_id ASC;



-- name: CreateGroupUser :execresult
INSERT INTO `groups_users` (gu_name, gu_description
) VALUES (?, ?); 

-- name: UpdateGroupUser :execrows
UPDATE `groups_users` SET gu_name = ?, gu_description = ? WHERE gu_id = ?;


-- name: DeleteGroupUser :execrows
DELETE FROM `groups_users`
WHERE gu_id = ?;
