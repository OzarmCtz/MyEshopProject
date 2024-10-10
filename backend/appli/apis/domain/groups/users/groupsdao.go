package users

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListGroupsUsers() ([]adm.GroupsUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	groups, err := adm.QueriesDb.ListGroupsUsers(ctx)
	return groups, err
}

func ListGroupsUsersByUser(userId int32) ([]adm.GroupsUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	groups, err := adm.QueriesDb.ListGroupsUserByUser(ctx, userId)
	return groups, err
}

func GetGroupUser(groupId int32) (adm.GroupsUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	group, err := adm.QueriesDb.GetGroupUser(ctx, groupId)
	return group, err
}

func GetGroupUserByName(groupName string) (adm.GroupsUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	group, err := adm.QueriesDb.GetGroupByName(ctx, groupName)
	return group, err
}

func InsertGroupUser(groupParams adm.CreateGroupUserParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateGroupUser(ctx, groupParams)
	return res, err
}

func UpdateGroupUser(groupParams adm.UpdateGroupUserParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateGroupUser(ctx, groupParams)
	return rows, err
}

func DeleteGroupUser(groupId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteGroupUser(ctx, groupId)
	return rows, err
}
