package settings

import (
	"fmt"
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasas "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/app/settings"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListAppSettingsPublic(c *gin.Context) {
	appSettings, err := aasas.SettingService.ListAppSettingsPublic()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(appSettings) == 0 {
		c.JSON(http.StatusNotFound, appSettings)
		return
	}

	c.JSON(http.StatusOK, appSettings)
}

func GetAppSettingsByKeyPublic(c *gin.Context) {
	appKey := c.Param("appKey")

	if appKey != "" {

		appSetting, err := aasas.SettingService.GetAppSettingsByKeyPublic(appKey)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, appSetting)

	}
	if appKey == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

// Private

func ListAppSettingPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		appSettings, err := aasas.SettingService.ListAppSettingsPrivate(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(appSettings) == 0 {
			c.JSON(http.StatusNotFound, appSettings)
			return
		}

		c.JSON(http.StatusOK, appSettings)
		return
	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func GetAppSettingsPrivate(c *gin.Context) {
	appSettingsId, appSettingsExists := c.Get(adg.APP_SETTINGS_ID)
	userLogin, userLoginExist := c.Get("userLogin")

	if appSettingsExists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		settingsId := appSettingsId.(int32)

		appSetting, err := aasas.SettingService.GetAppSettingsPrivate(currentUser, settingsId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, appSetting)
		return
	}

	if !appSettingsExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func InsertAppSettingsPrivate(c *gin.Context) {
	userLogin, userLoginExist := c.Get("userLogin")

	if userLoginExist {
		currentUsser := userLogin.(aadu.AppliUserLogin)

		appSettingParams := adm.CreateAppSettingsParams{}
		if err := c.ShouldBindJSON(&appSettingParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res, err := aasas.SettingService.InsertAppSettingsPrivate(currentUsser, appSettingParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func UpdateAppSettingPrivate(c *gin.Context) {
	appSettingsId, appSettingsExists := c.Get(adg.APP_SETTINGS_ID)
	userLogin, userLoginExist := c.Get("userLogin")

	if appSettingsExists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		settingsId := appSettingsId.(int32)
		updateMysqlSettings := adm.UpdateAppSettingsParams{AsID: settingsId}
		if err := c.ShouldBindJSON(&updateMysqlSettings); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasas.SettingService.UpdateAppSettingsPrivate(currentUser, settingsId, updateMysqlSettings)
		isErr := aac.EditErrorResponse(rows, settingsId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("settings %d updated ", settingsId))
		}
		return
	}

	if !appSettingsExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func DeleteAppSettingPrivate(c *gin.Context) {
	appSettingsId, appSettingsExists := c.Get(adg.APP_SETTINGS_ID)
	userLogin, userLoginExist := c.Get("userLogin")

	if appSettingsExists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		settingsId := appSettingsId.(int32)

		rows, err := aasas.SettingService.DeleteAppSettingsPrivate(currentUser, settingsId)
		isErr := aac.EditErrorResponse(rows, settingsId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("settings %d deleted ", settingsId))
		}
		return
	}

	if !appSettingsExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}
