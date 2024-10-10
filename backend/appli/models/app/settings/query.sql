-- name: GetAppSettings :one
SELECT * FROM app_settings
WHERE as_id = ? LIMIT 1;


-- name: GetAppSettingsByKey :one
SELECT * FROM app_settings
WHERE as_key = ?  LIMIT 1;

-- name: ListAppSettings :many
SELECT * FROM app_settings
ORDER BY as_id ASC;

-- name: CreateAppSettings :execresult
INSERT INTO app_settings (as_key, as_value ,as_description, as_last_updated  
) VALUES (?, ?, ?, ?);

-- name: DeleteAppSettings :execrows
DELETE FROM app_settings
WHERE as_id = ?;

-- name: UpdateAppSettings :execrows
UPDATE app_settings SET as_key = ?,as_value = ?, as_description = ?, as_last_updated = ? WHERE as_id = ?;