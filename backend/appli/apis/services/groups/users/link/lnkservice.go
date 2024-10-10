package linkservice

import (
	"errors"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadg "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadgul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users/link"
	aadm "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aua "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth"
)

var (
	GroupsUsersLinkService groupsUsersLinkServiceInterface = &groupsUsersLink{}
)

type groupsUsersLink struct{}

type groupsUsersLinkServiceInterface interface {
	InsertGroupsUsersLinkByGroupName(params adm.CreateGroupUserLinkByGroupNameParams) (adm.GroupsUsersLink, error)
	UpdateGroupUserLink(currentUser aadm.AppliUserLogin, groupUserParams adm.UpdateGroupUserLinkParams) (int64, error)
	GetGroupUserLinkByUser(gulUserId int32) (adm.GroupsUsersLink, error)
}

func (ls *groupsUsersLink) InsertGroupsUsersLinkByGroupName(params adm.CreateGroupUserLinkByGroupNameParams) (adm.GroupsUsersLink, error) {
	var groupsUsersLink adm.GroupsUsersLink

	res, err := aadgul.InsertGroupUserLinkByGroupName(params)
	if err != nil {
		return groupsUsersLink, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return groupsUsersLink, err
	}

	groupsUsersLink, err = aadgul.GetGroupUserLink(int32(id))
	if err != nil {
		return groupsUsersLink, err
	}

	return groupsUsersLink, nil
}

func (ls *groupsUsersLink) UpdateGroupUserLink(currentUser aadm.AppliUserLogin, groupUserParams adm.UpdateGroupUserLinkParams) (int64, error) {
	userId := groupUserParams.GulUserID
	groupId := groupUserParams.GulGroupID

	groupName, err := getGroupName(groupId)
	if err != nil {
		return 0, err
	}

	// if is super admin

	isSuperAdmin, err := aua.IsRealySuperAdmin(currentUser.User.UID)
	if err != nil {
		return 0, err
	}
	if isSuperAdmin {

		// if user is not the same as the current user
		if currentUser.User.UID != userId {
			isSuperUserAdmin, err := isUserStatusChecked(userId, adg.SUPER_ADMIN_STATUS)
			if err != nil {
				return 0, err
			}
			// if user to update is not super admin
			if !isSuperUserAdmin {
				return aadgul.UpdateGroupUsersLink(groupUserParams)
			}
		} else {
			return aadgul.UpdateGroupUsersLink(groupUserParams)
		}
	}

	isAdmin, err := aua.IsRealyAdmin(currentUser.User.UID)
	if err != nil {
		return 0, err
	}

	// if is admin
	if isAdmin {

		userToUpdateIsSuperAdmin, err := aua.IsRealySuperAdmin(groupUserParams.GulID)
		if err != nil {
			return 0, errors.New("an admin can't update superadmin")
		}

		if userToUpdateIsSuperAdmin {
			return 0, errors.New("an admin user can't add user into the super admin group")
		}
		// if user is not the same as the current user
		if currentUser.User.UID != userId {
			isUserAdmin, err := isUserStatusChecked(userId, adg.ADMIN_STATUS)
			if err != nil {
				return 0, err
			}

			// if user to update is not admin or superadmin
			if !isUserAdmin {
				if groupName == adg.SUPER_ADMIN_STATUS || groupName == adg.ADMIN_STATUS {
					return 0, errors.New("an admin user can't add user into the admin or super admin group")
				}
				return aadgul.UpdateGroupUsersLink(groupUserParams)
			}
		} else {
			if groupName == adg.SUPER_ADMIN_STATUS {
				return 0, errors.New("an admin user can't add himself into the super admin group")
			}
			return aadgul.UpdateGroupUsersLink(groupUserParams)
		}
	}
	err = errors.New("user doesn't any have privileges to update group user")
	return 0, err
}

func (ls *groupsUsersLink) GetGroupUserLinkByUser(gulUserId int32) (adm.GroupsUsersLink, error) {
	return aadgul.GetGroupUserLinkByUser(gulUserId)
}

// Utility functions

func isUserStatusChecked(userId int32, status string) (bool, error) {
	groups, err := aadg.ListGroupsUsersByUser(userId)
	if err != nil {
		return false, err
	}
	return aua.CheckUserStatus(groups, status), nil
}

func getGroupName(groupId int32) (string, error) {
	group, err := aadg.GetGroupUser(groupId)
	if err != nil {
		return "", err
	}
	return group.GuName, nil
}
