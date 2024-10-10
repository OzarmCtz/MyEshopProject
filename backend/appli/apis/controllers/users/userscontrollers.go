package userscontroller

import (
	"fmt"
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasgul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users/link"
	aasu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		users, err := aasu.UsersService.ListUsers(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(users) == 0 {
			c.JSON(http.StatusNotFound, users)
			return
		}

		c.JSON(http.StatusOK, users)
		return

	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

}

func Get(c *gin.Context) {
	usersId, usersIdExists := c.Get(adg.USER_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if usersIdExists && userLoginExist {
		userId := usersId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		user, err := aasu.UsersService.GetUser(currentUser, userId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
		return
	}
	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !usersIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func Delete(c *gin.Context) {
	usersId, userIdExists := c.Get(adg.USER_ID)
	userLogin, userExists := c.Get("userLogin")

	if userIdExists && userExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		uId := usersId.(int32)
		rows, err := aasu.UsersService.DeleteUser(currentUser, uId)
		isErr := aac.EditErrorResponse(rows, uId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("user %d deleted ", uId))
		}
		return
	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func Update(c *gin.Context) {
	usersId, usersIdexists := c.Get(adg.USER_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if usersIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		uId := usersId.(int32)
		updateMysqlUser := adm.UpdateUserParams{UID: uId}
		if err := c.ShouldBindJSON(&updateMysqlUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasu.UsersService.UpdateUser(currentUser, updateMysqlUser)
		isErr := aac.EditErrorResponse(rows, uId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("user %d updated ", uId))
		}
		return
	}
	if !usersIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func SignUpUser(c *gin.Context) {
	createMysqlUser := adm.CreateUserParams{}

	if err := c.ShouldBindJSON(&createMysqlUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//tokenFirebase, err := af.CreateUserInFb(createMysqlUser)
	//if err != nil {
	//c.JSON(400, gin.H{"error": err.Error()})
	//return
	//}

	user, err := aasu.UsersService.InsertUser(createMysqlUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createGroupsUsersLink := adm.CreateGroupUserLinkByGroupNameParams{}
	createGroupsUsersLink.GulUserID = user.UID
	createGroupsUsersLink.GuName = adg.CLIENT_STATUS

	_, err = aasgul.GroupsUsersLinkService.InsertGroupsUsersLinkByGroupName(createGroupsUsersLink)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		//"token":  tokenFirebase,
		"userId": user.UID,
	}
	c.JSON(201, &response)
}
