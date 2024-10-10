-- name: GetGroupPrivilege :one
SELECT * FROM `groups_privileges`
WHERE gp_id = ? LIMIT 1;

-- name: GetGroupPrivilegesByUserId :many
SELECT DISTINCT p.* FROM groups_privileges p
INNER JOIN groups_privileges_link gp ON p.gp_id = gp.gpl_groups_privileges_id
INNER JOIN groups_users_link gul ON gp.gpl_groups_users_id = gul.gul_group_id 
INNER JOIN users u ON gul.gul_user_id = u.u_id
WHERE u.u_id = ? ORDER BY p.gp_id ASC;


-- name: ListGroupsPrivileges :many
SELECT * FROM `groups_privileges`
ORDER BY gp_id ASC;

-- name: CreateGroupPrivilege :execresult
INSERT INTO `groups_privileges` (gp_path
) VALUES (?);

-- name: DeleteGroupPrivilege :execrows
DELETE FROM `groups_privileges`
WHERE gp_id = ?;

-- name: UpdateGroupPrivilege :execrows
UPDATE `groups_privileges` SET gp_path = ? 
WHERE gp_id = ?;