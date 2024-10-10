package settings

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListAppSettings() ([]adm.AppSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	settings, err := adm.QueriesDb.ListAppSettings(ctx)
	return settings, err
}

func GetAppSettings(settingsId int32) (adm.AppSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	settings, err := adm.QueriesDb.GetAppSettings(ctx, settingsId)
	return settings, err
}

func GetAppSettingsByKey(appKey string) (adm.AppSetting, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	settings, err := adm.QueriesDb.GetAppSettingsByKey(ctx, appKey)
	return settings, err
}

func InsertAppSettings(settingsParams adm.CreateAppSettingsParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.CreateAppSettings(ctx, settingsParams)
	return rows, err
}

func DeleteAppSettings(settingsId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteAppSettings(ctx, settingsId)
	return rows, err
}

func UpdateAppSettings(settingsId int32, settingsParams adm.UpdateAppSettingsParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateAppSettings(ctx, settingsParams)
	return rows, err
}
