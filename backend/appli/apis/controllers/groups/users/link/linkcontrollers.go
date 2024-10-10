package linkcontrollers

import (
	"fmt"
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users"
	aasgul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users/link"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	gulIdParam, gulIdParamExists := c.Get(adg.GROUP_USER_LINK_ID)

	if gulIdParamExists {
		gulId := gulIdParam.(int32)
		gul, err := aasgul.GroupsUsersLinkService.GetGroupUserLinkByUser(gulId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		gu, err := aasgu.GroupsUsersService.GetGroupUsers(gul.GulGroupID)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, &gu)
		return
	}
	if !gulIdParamExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func Update(c *gin.Context) {
	userLogin, userExists := c.Get("userLogin")
	gulIdParam, guIdParamExists := c.Get(adg.GROUP_USER_LINK_ID)

	if userExists && guIdParamExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		gulID := gulIdParam.(int32)
		updateGroupUserLnkParams := adm.UpdateGroupUserLinkParams{}
		updateGroupUserLnkParams.GulID = gulID
		if err := c.ShouldBindJSON(&updateGroupUserLnkParams); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rows, err := aasgul.GroupsUsersLinkService.UpdateGroupUserLink(currentUser, updateGroupUserLnkParams)
		isErr := aac.EditErrorResponse(rows, gulID, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("groups users %d updated ", gulID))
		}
		return
	}
	if !guIdParamExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}
