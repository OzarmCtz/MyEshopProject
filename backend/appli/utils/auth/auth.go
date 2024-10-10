package authutils

import (
	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadgul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users/link"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func checkUserStatus(user aadu.AppliUserLogin, status string) bool {
	for _, group := range user.Groups {
		if group.GuName == status {
			return true
		}
	}
	return false
}

func CheckUserStatus(groups []adm.GroupsUser, status string) bool {
	for _, group := range groups {
		if group.GuName == status {
			return true
		}
	}
	return false
}

func IsUserSuperAdmin(user aadu.AppliUserLogin) bool {
	return checkUserStatus(user, adg.SUPER_ADMIN_STATUS)
}

func IsUserAdminStatus(user aadu.AppliUserLogin) bool {
	return checkUserStatus(user, adg.ADMIN_STATUS)
}

func IsUserAdmin(user aadu.AppliUserLogin) bool {
	return IsUserSuperAdmin(user) || checkUserStatus(user, adg.ADMIN_STATUS)
}

func IsRealySuperAdmin(uId int32) (bool, error) {
	groupsUserLink, err := aadgul.GetGroupUserLinkByUser(uId)
	if err != nil {
		return false, err
	}

	groupsUsersName, err := aadgu.GetGroupUser(groupsUserLink.GulGroupID)
	if err != nil {
		return false, err
	}

	return groupsUsersName.GuName == adg.SUPER_ADMIN_STATUS, nil

}

func IsRealyAdmin(uId int32) (bool, error) {
	groupsUserLink, err := aadgul.GetGroupUserLinkByUser(uId)
	if err != nil {
		return false, err
	}

	groupsUsersName, err := aadgu.GetGroupUser(groupsUserLink.GulGroupID)
	if err != nil {
		return false, err
	}

	return groupsUsersName.GuName == adg.ADMIN_STATUS, nil

}

func FindPrivilege(privileges []adm.GroupsPrivilege, privilege string) bool {
	for _, p := range privileges {
		if p.GpPath.String == privilege {
			return true
		}
	}
	return false
}
