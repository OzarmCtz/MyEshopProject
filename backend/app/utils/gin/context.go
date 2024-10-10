package context

import (
	"errors"
	"fmt"
	"time"

	"firebase.google.com/go/auth"
	af "github.com/OzarmCtz/e_shop_backend_v1/app/firebase"
	aulz "github.com/OzarmCtz/e_shop_backend_v1/app/utils/logger/zap"
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	aadgp "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/privileges"
	aadg "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

type User struct {
	AppliUser *adm.User
}

func GetCentralUserFromContext(token string) (adm.User, error) {
	var user adm.User
	userRecord, err := GetUserRecordFromContext(token)
	if err != nil {
		return user, err
	}

	// Get user in db by mail
	filteredUser, err := aasu.UsersService.GetUserByEmail(userRecord.Email)
	if err != nil {
		return user, err
	}

	user = adm.User(aasu.SetFilteredUserToUser(filteredUser))
	return user, nil
}

func GetAplliUserFromContext(currentUser aadu.AppliUserLogin, token string) (adm.User, error) {
	var user adm.User
	userRecord, err := GetUserRecordFromContext(token)
	if err != nil {
		return user, err
	}

	// Get user in db by mail
	filteredUser, err := aasu.UsersService.GetUserByEmail(userRecord.Email)
	if err != nil {
		return user, err
	}

	user = adm.User(aasu.SetFilteredUserToUser(filteredUser))
	return user, nil
}
func GetUserRecordFromContext(token string) (*auth.UserRecord, error) {
	var userRecord *auth.UserRecord

	authToken, err := af.GetAuthToken(token)
	if err != nil {
		return userRecord, err
	}

	userRecord, err = af.AuthByUid(authToken.UID)
	if err != nil {
		return userRecord, err
	}
	return userRecord, nil
}

func SaveAppliUserLoginFromContext(c *gin.Context, appliUser adm.User) error {
	user := User{
		AppliUser: &appliUser,
	}
	return SaveUserLoginFromContext(c, user)
}

func SaveUserLoginFromContext(c *gin.Context, user User) error {
	var appliGroupsUser []adm.GroupsUser
	var userId int32
	var err error

	appliGroupsUser, err = aadg.ListGroupsUsers()
	userId = user.AppliUser.UID
	if err != nil {
		aulz.Info(fmt.Sprintf("error in getting groups %v for user %d", err, userId))
		return err
	}

	privileges, err := aadgp.GetGroupPrivilegesByUserId(userId)
	if err != nil {
		aulz.Info(fmt.Sprintf("error in getting privileges %v for user %d", err, userId))
		return err
	}

	userLogin := aadu.AppliUserLogin{
		User:       *user.AppliUser,
		LoginTime:  time.Now().Format("2006-01-02 15:04:05"),
		Privileges: privileges,
		Groups:     appliGroupsUser,
	}
	c.Set("userLogin", userLogin)

	return nil
}

func GetParamId(c *gin.Context, paramId string) (int32, error) {
	idParam := c.Param(paramId)
	if idParam == "" {
		return 0, nil
	}
	id, err := aus.StrToInt64(idParam)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return int32(id), nil
}
