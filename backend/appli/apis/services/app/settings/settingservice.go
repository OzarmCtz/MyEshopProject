package settings

import (
	"time"

	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	aadas "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/app/settings"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	SettingService settingServiceInterface = &settingService{}
)

type AppSetting []adm.AppSetting
type AppSettings adm.AppSetting
type settingService struct{}

type settingServiceInterface interface {
	GetAppSettingsPrivate(currentUser adu.AppliUserLogin, appSettingsId int32) (AppSettingPvResponse, error)
	ListAppSettingsPrivate(currentUser adu.AppliUserLogin) ([]PrivateAppSettingResponse, error)
	InsertAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsParams adm.CreateAppSettingsParams) (AppSettingPvResponse, error)
	UpdateAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsId int32, settingsParams adm.UpdateAppSettingsParams) (int64, error)
	DeleteAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsId int32) (int64, error)
	GetAppSettingsByKeyPublic(appKey string) (AppSettingPbResponse, error)
	ListAppSettingsPublic() ([]PublicAppSettingResponse, error)
}

// PUBLIC

func (ss *settingService) GetAppSettingsByKeyPublic(appKey string) (AppSettingPbResponse, error) {
	appSettingRes, err := aadas.GetAppSettingsByKey(appKey)
	if err != nil {
		return AppSettingPbResponse{}, err
	}

	publicAppSetting := PublicAppSettingResponse{
		AsKey:   appSettingRes.AsKey,
		AsValue: appSettingRes.AsValue,
	}

	return AppSettingPbResponse(publicAppSetting), err
}

func (ss *settingService) ListAppSettingsPublic() ([]PublicAppSettingResponse, error) {
	rawAppSettings, err := aadas.ListAppSettings()
	if err != nil {
		return nil, err
	}

	publicAppSettings := make([]PublicAppSettingResponse, 0, len(rawAppSettings))
	for _, appSetting := range rawAppSettings {
		publicAppSetting := PublicAppSettingResponse{
			AsKey:   appSetting.AsKey,
			AsValue: appSetting.AsValue,
		}
		publicAppSettings = append(publicAppSettings, publicAppSetting)
	}

	return publicAppSettings, err
}

// PRIVATE

func (ss *settingService) GetAppSettingsPrivate(currentUser adu.AppliUserLogin, appSettingsId int32) (AppSettingPvResponse, error) {
	appSettingRes, err := aadas.GetAppSettings(appSettingsId)
	if err != nil {
		return AppSettingPvResponse{}, err
	}

	privateAppSetting := PrivateAppSettingResponse{
		AsID:          appSettingRes.AsID,
		AsKey:         appSettingRes.AsKey,
		AsValue:       appSettingRes.AsValue,
		AsDescription: aus.NullStringToString(appSettingRes.AsDescription),
		AsLastUpdated: appSettingRes.AsLastUpdated,
	}

	return AppSettingPvResponse(privateAppSetting), err
}

func (ss *settingService) ListAppSettingsPrivate(currentUser adu.AppliUserLogin) ([]PrivateAppSettingResponse, error) {
	rawAppSettings, err := aadas.ListAppSettings()
	if err != nil {
		return nil, err
	}

	privateAppSettings := make([]PrivateAppSettingResponse, 0, len(rawAppSettings))
	for _, appSetting := range rawAppSettings {
		privateAppSetting := PrivateAppSettingResponse{
			AsID:          appSetting.AsID,
			AsKey:         appSetting.AsKey,
			AsValue:       appSetting.AsValue,
			AsDescription: aus.NullStringToString(appSetting.AsDescription),
			AsLastUpdated: appSetting.AsLastUpdated,
		}
		privateAppSettings = append(privateAppSettings, privateAppSetting)
	}

	return privateAppSettings, err
}

func (ss *settingService) InsertAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsParams adm.CreateAppSettingsParams) (AppSettingPvResponse, error) {
	settingsParams.AsLastUpdated = time.Now()

	res, err := aadas.InsertAppSettings(settingsParams)
	if err != nil {
		return AppSettingPvResponse{}, err
	}

	appSettingId, err := res.LastInsertId()
	if err != nil {
		return AppSettingPvResponse{}, err
	}

	appSettingRes, err := aadas.GetAppSettings(int32(appSettingId))
	if err != nil {
		return AppSettingPvResponse{}, err
	}

	privateAppSetting := PrivateAppSettingResponse{
		AsID:          appSettingRes.AsID,
		AsKey:         appSettingRes.AsKey,
		AsValue:       appSettingRes.AsValue,
		AsDescription: aus.NullStringToString(appSettingRes.AsDescription),
		AsLastUpdated: appSettingRes.AsLastUpdated,
	}

	return AppSettingPvResponse(privateAppSetting), err
}

func (ss *settingService) UpdateAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsId int32, settingsParams adm.UpdateAppSettingsParams) (int64, error) {
	settingsParams.AsLastUpdated = time.Now()

	rows, err := aadas.UpdateAppSettings(settingsId, settingsParams)
	if err != nil {
		return 0, err
	}

	return rows, err
}

func (ss *settingService) DeleteAppSettingsPrivate(currentUser adu.AppliUserLogin, settingsId int32) (int64, error) {
	rows, err := aadas.DeleteAppSettings(settingsId)
	if err != nil {
		return 0, err
	}

	return rows, err
}
