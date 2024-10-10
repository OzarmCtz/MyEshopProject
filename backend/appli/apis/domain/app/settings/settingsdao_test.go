package settings

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestAppSettingDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	insertAppSettingsparams := adm.CreateAppSettingsParams{
		AsKey:         "TestKey",
		AsValue:       "TestValue",
		AsDescription: adm.NullString{NullString: sql.NullString{String: "Test Description", Valid: true}},
		AsLastUpdated: time.Now(),
	}

	resInsertAppSetting, err := InsertAppSettings(insertAppSettingsparams)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	appSettingId, err := resInsertAppSetting.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, appSettingId, int64(1))

	getAppSettings, err := GetAppSettings(int32(appSettingId))

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, getAppSettings.AsKey, "TestKey")

	getAppSettingsByKey, err := GetAppSettingsByKey("TestKey")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, getAppSettingsByKey.AsDescription, adm.NullString{NullString: sql.NullString{String: "Test Description", Valid: true}})

	updateAppSettingsParams := adm.UpdateAppSettingsParams{
		AsKey:         "TestKey",
		AsValue:       "TestValue Updated",
		AsDescription: adm.NullString{NullString: sql.NullString{String: "Test Description Updated", Valid: true}},
		AsLastUpdated: time.Now(),
		AsID:          int32(appSettingId),
	}

	rows, err := UpdateAppSettings(int32(appSettingId), updateAppSettingsParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	appSettings, err := ListAppSettings()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, appSettings[0].AsValue, "TestValue Updated")

	rows, err = DeleteAppSettings(int32(appSettingId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))
}
