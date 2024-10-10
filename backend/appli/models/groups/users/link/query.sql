-- name: GetGroupUserLink :one
SELECT * FROM `groups_users_link`
WHERE gul_id = ? LIMIT 1;

-- name: GetGroupUserLinkByUser :one
SELECT * FROM `groups_users_link`
WHERE gul_user_id = ? LIMIT 1;


-- name: ListGroupsUsersLink :many
SELECT * FROM `groups_users_link`
ORDER BY gul_id ASC;

-- name: CreateGroupUserLink :execresult
INSERT INTO `groups_users_link` (gul_user_id , gul_group_id
) VALUES (? , ?);

-- name: CreateGroupUserLinkByGroupName :execresult
INSERT INTO `groups_users_link` (gul_user_id, gul_group_id)
SELECT ?, gu.gu_id FROM `groups_users` gu
WHERE gu.gu_name = ?;

-- name: DeleteGroupUserLink :execrows
DELETE FROM `groups_users_link`
WHERE gul_id = ?;

-- name: UpdateGroupUserLink :execrows
UPDATE `groups_users_link` SET gul_user_id = ? , gul_group_id = ?
WHERE gul_id = ?;