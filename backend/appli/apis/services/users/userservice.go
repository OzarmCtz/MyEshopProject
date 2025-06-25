package userservice

import (
	"errors"
	"time"

	//	"firebase.google.com/go/auth"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	af "github.com/OzarmCtz/e_shop_backend_v1/app/firebase"
	aadg "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users/link"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users"
	aasgul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/groups/users/link"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aua "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth"
	auad "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth/data"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type Users []adm.User
type User adm.User
type FilteredUser User
type FilteredUsers []FilteredUser
type usersService struct{}

type usersServiceInterface interface {
	ListUsers(currentUser adu.AppliUserLogin) ([]UserInfo, error)
	GetUser(currentUser adu.AppliUserLogin, userId int32) (adm.User, error)
	GetUserByEmail(email string) (FilteredUser, error)
	UpdateUser(currentUser adu.AppliUserLogin, userParams adm.UpdateUserParams) (int64, error)
	DeleteUser(currentUser adu.AppliUserLogin, userID int32) (int64, error)
	InsertUser(userParams adm.CreateUserParams) (adm.User, error)
}

func (us *usersService) InsertUser(userParams adm.CreateUserParams) (adm.User, error) {
	var user adm.User

	userParams.URegisterDate = time.Now()
	res, err := adu.InsertUser(userParams)
	if err != nil {
		return user, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}

	user, err = adu.GetUser(int32(id))
	if err != nil {
		return user, err
	}

	return user, nil

}

func (us *usersService) GetUser(currentUser adu.AppliUserLogin, userId int32) (adm.User, error) {
	if userId != currentUser.User.UID {

		userGroup, err := aadgu.GetGroupUserLinkByUser(currentUser.User.UID)
		if err != nil {
			return adm.User{}, err
		}

		userGroupName, err := aadg.GetGroupUser(userGroup.GulGroupID)
		if err != nil {
			return adm.User{}, err
		}

		if userGroupName.GuName == global.CLIENT_STATUS {
			return adm.User{}, errors.New("you are not authorized to Get this user")
		}
	}
	userRes, err := adu.GetUser(userId)
	if err != nil {
		return adm.User{}, err
	}

	return userRes, nil
}

func (us *usersService) UpdateUser(currentUser adu.AppliUserLogin, userParams adm.UpdateUserParams) (int64, error) {

	if userParams.UID != currentUser.User.UID {
		userGroup, err := aadgu.GetGroupUserLinkByUser(currentUser.User.UID)
		if err != nil {
			return -1, err
		}

		userGroupName, err := aadg.GetGroupUser(userGroup.GulGroupID)
		if err != nil {
			return -1, err
		}

		if userGroupName.GuName == global.CLIENT_STATUS || userGroupName.GuName == global.ADMIN_STATUS {
			return 0, errors.New("you are not authorized to update this user")
		}

		userToUpdateGroup, err := aadgu.GetGroupUserLinkByUser(userParams.UID)
		if err != nil {
			return -1, err
		}

		userToUpdateGroupName, err := aadg.GetGroupUser(userToUpdateGroup.GulGroupID)
		if err != nil {
			return -1, err
		}

		if userGroupName.GuName == global.SUPER_ADMIN_STATUS && userToUpdateGroupName.GuName == global.SUPER_ADMIN_STATUS {
			return 0, errors.New("you are not authorized to update this user")
		}

	}

	fieldsErrorMessage, err := auad.CheckUpdateUserParams(userParams)
	if err != nil {
		return -1, err
	} else if len(fieldsErrorMessage) > 0 {
		for _, errMsg := range fieldsErrorMessage {
			return -1, errors.New(errMsg)
		}
	}

	user, err := us.GetUser(currentUser, userParams.UID)
	if err != nil {
		return -1, err
	}
	userParams.URegisterDate = user.URegisterDate
	userParams.UID = user.UID

	updateUserInDb, err := adu.UpdateUser(userParams)
	if err != nil {
		return -1, err
	}

	/*if userParams.UEmail != user.UEmail || userParams.UID != user.UID || userParams.UIsDisabled != user.UIsDisabled {
		userFirebaseUid, err := af.GetUIDFromEmail(user.UEmail)
		if err != nil {
			return -1, err
		}

		err = af.UpdateUserEmailAndPassword(userFirebaseUid, userParams.UEmail, userParams.UID, userParams.UIsDisabled)
		if err != nil {
			return -1, err
		}
	}*/

	return updateUserInDb, nil

}

func (us *usersService) DeleteUser(currentUser adu.AppliUserLogin, userID int32) (int64, error) {
	if userID != currentUser.User.UID {
		userGroup, err := aadgu.GetGroupUserLinkByUser(currentUser.User.UID)
		if err != nil {
			return -1, err
		}

		userGroupName, err := aadg.GetGroupUser(userGroup.GulGroupID)
		if err != nil {
			return -1, err
		}

		if userGroupName.GuName == global.CLIENT_STATUS {
			return 0, errors.New("you are not authorized to delete this user")
		}

		userToDeleteGroup, err := aadgu.GetGroupUserLinkByUser(userID)
		if err != nil {
			return -1, err
		}

		userToDeleteGroupName, err := aadg.GetGroupUser(userToDeleteGroup.GulGroupID)
		if err != nil {
			return -1, err
		}

		if (userGroupName.GuName == global.SUPER_ADMIN_STATUS && userToDeleteGroupName.GuName == global.SUPER_ADMIN_STATUS) ||
			(userGroupName.GuName == global.ADMIN_STATUS && userToDeleteGroupName.GuName == global.ADMIN_STATUS ||
				userToDeleteGroupName.GuName == global.SUPER_ADMIN_STATUS) {
			return 0, errors.New("you are not authorized to delete this user")
		}
	}

	userEmail, err := us.GetUser(currentUser, userID)
	if err != nil {
		return -1, err
	}

	userFirebaseUid, err := af.GetUIDFromEmail(userEmail.UEmail)
	if err != nil {
		return -1, err
	}
	err = af.DeleteUser(userFirebaseUid)
	if err != nil {
		return -1, err
	}

	return adu.DeleteUser(userID)
}

func (us *usersService) GetUserByEmail(email string) (FilteredUser, error) {
	user, err := adu.GetUserByEmail(email)
	if err != nil {
		return FilteredUser{}, err
	}
	return FilteredUser(user), nil

}

func (us *usersService) ListUsers(currentUser adu.AppliUserLogin) ([]UserInfo, error) {
	var userInfos []UserInfo

	userGroup, err := aadgu.GetGroupUserLinkByUser(currentUser.User.UID)
	if err != nil {
		return nil, err
	}

	userGroupName, err := aadg.GetGroupUser(userGroup.GulGroupID)
	if err != nil {
		return nil, err
	}

	if userGroupName.GuName == global.CLIENT_STATUS {
		return nil, errors.New("you are not authorized to list users")
	}

	userList, err := adu.ListUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range userList {
		gul, err := aasgul.GroupsUsersLinkService.GetGroupUserLinkByUser(user.UID)
		if err != nil {
			continue // or handle error depending on your requirement
		}
		gu, err := aasgu.GroupsUsersService.GetGroupUsers(gul.GulGroupID)
		if err != nil {
			continue // or handle error depending on your requirement
		}

		userInfo := UserInfo{
			UID:           user.UID,
			UUid:          user.UUid,
			UEmail:        user.UEmail,
			URegisterDate: user.URegisterDate,
			UIsDisabled:   user.UIsDisabled,
			UserStatus:    gu.GuName,
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

// Utility functions
func SetFilteredUserToUser(fu FilteredUser) User {
	// TODO : FIX THIS WARNING
	return User{
		UID:           fu.UID,
		UEmail:        fu.UEmail,
		URegisterDate: fu.URegisterDate,
		UIsDisabled:   fu.UIsDisabled,
	}
}

func (users Users) Marshal(currentUser adu.AppliUserLogin) FilteredUsers {
	var filteredUsers FilteredUsers

	for _, user := range users {
		newUser := User{
			UID:           user.UID,
			UEmail:        user.UEmail,
			URegisterDate: user.URegisterDate,
			UIsDisabled:   user.UIsDisabled,
		}
		filteredUser := newUser.Marshal(currentUser)
		filteredUsers = append(filteredUsers, filteredUser)
	}
	return filteredUsers
}

func (user User) Marshal(currentUser adu.AppliUserLogin) FilteredUser {

	if aua.IsUserAdmin(currentUser) || user.UID == currentUser.User.UID {
		return FilteredUser(user)
	}

	newUser := FilteredUser{}
	newUser.UID = user.UID

	return newUser
}

type UserInfo struct {
	UID           int32     `json:"u_id"`
	UUid          string    `json:"u_uid"`
	UEmail        string    `json:"u_email"`
	URegisterDate time.Time `json:"u_register_date"`
	UIsDisabled   bool      `json:"u_is_disabled"`
	UserStatus    string    `json:"user_status"` // Ajout de UserStatus dans la structure
}
