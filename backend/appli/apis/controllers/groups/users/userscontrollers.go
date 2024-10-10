package userscontrollers

import (
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users"
	"github.com/gin-gonic/gin"
)

func ListGroupUserPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		usersGroups, err := aasgu.GroupsUsersService.ListGroupUser(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(usersGroups) == 0 {
			c.JSON(http.StatusNotFound, usersGroups)
			return
		}

		c.JSON(http.StatusOK, usersGroups)
		return
	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}
