package userservice

import (
	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	aadm "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	GroupsUsersService groupsUsersServiceInterface = &groupsUsersService{}
)

type groupsUsersService struct{}

type groupsUsersServiceInterface interface {
	GetGroupUsers(groupId int32) (adm.GroupsUser, error)
	ListGroupUser(currentUser aadm.AppliUserLogin) ([]adm.GroupsUser, error)
}

func (cs *groupsUsersService) GetGroupUsers(groupId int32) (adm.GroupsUser, error) {
	return aadgu.GetGroupUser(groupId)
}

func (cs *groupsUsersService) ListGroupUser(currentUser aadm.AppliUserLogin) ([]adm.GroupsUser, error) {
	return aadgu.ListGroupsUsers()
}
