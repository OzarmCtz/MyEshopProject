-- name: GetGroupPrivilegeLink :one
SELECT * FROM `groups_privileges_link`
WHERE gpl_id = ? LIMIT 1;

-- name: ListGroupPrivilegesLink :many
SELECT * FROM `groups_privileges_link`
ORDER BY gpl_id ASC;

-- name: CreateGroupPrivilegeLink :execresult
INSERT INTO `groups_privileges_link` (gpl_groups_users_id , gpl_groups_privileges_id
) VALUES ( 
  ? , ?
);

-- name: UpdateGroupPrivilegeLink :execrows
UPDATE `groups_privileges_link` SET gpl_groups_users_id = ? , gpl_groups_privileges_id = ?
WHERE gpl_id = ?;

-- name: DeleteGroupPrivilegeLink :execrows
DELETE FROM `groups_privileges_link`
WHERE gpl_id = ?;
