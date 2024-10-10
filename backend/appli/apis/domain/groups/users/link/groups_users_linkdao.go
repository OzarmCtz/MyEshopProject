package link

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	"github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adt "github.com/OzarmCtz/e_shop_backend_v1/appli/data/tests"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListGroupsUsersLink() ([]adm.GroupsUsersLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	groupsUsersLnk, err := adm.QueriesDb.ListGroupsUsersLink(ctx)
	return groupsUsersLnk, err
}

func GetGroupUserLink(guId int32) (adm.GroupsUsersLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	groupUserLnk, err := adm.QueriesDb.GetGroupUserLink(ctx, guId)
	return groupUserLnk, err
}

func GetGroupUserLinkByUser(gulUserId int32) (adm.GroupsUsersLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	groupUserLnk, err := adm.QueriesDb.GetGroupUserLinkByUser(ctx, gulUserId)
	return groupUserLnk, err
}

func InsertGroupUserLinkByGroupName(params adm.CreateGroupUserLinkByGroupNameParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateGroupUserLinkByGroupName(ctx, params)
	return res, err
}

func InsertGroupUsersLink(groupUserLnkParams adm.CreateGroupUserLinkParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateGroupUserLink(ctx, groupUserLnkParams)
	return res, err
}

func UpdateGroupUsersLink(groupUserLnkParams adm.UpdateGroupUserLinkParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateGroupUserLink(ctx, groupUserLnkParams)
	return rows, err
}

func DeleteGroupsUserLink(guId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteGroupUserLink(ctx, guId)
	return rows, err
}

// UTILITY FUNCTIONS FOR TESTS
type GroupsDependencies struct {
	User            *sql.Result
	Group           *sql.Result
	GroupUserLink   *sql.Result
	UserId          int64
	GroupUserId     int64
	GroupUserLinkId int64
}

func GetGroupsDependencies() (*GroupsDependencies, error) {
	userParams := adt.CreateUserParams

	resInsertUser, err := users.InsertUser(userParams)
	if err != nil {
		return nil, err
	}

	CreateGroupParams := adm.CreateGroupUserParams{
		GuName:        "TEST_STATUS",
		GuDescription: adm.NullString{NullString: sql.NullString{String: "Test Groups Status", Valid: true}},
	}

	resInsertGroupUser, err := aadgu.InsertGroupUser(CreateGroupParams)
	if err != nil {
		return nil, err
	}

	userId, err := resInsertUser.LastInsertId()
	if err != nil {
		return nil, err
	}

	groupUserId, err := resInsertGroupUser.LastInsertId()
	if err != nil {
		return nil, err
	}

	groupUserLinkCreateParams := adm.CreateGroupUserLinkParams{
		GulUserID:  int32(userId),
		GulGroupID: int32(groupUserId),
	}

	resInsertGroupUserLink, err := InsertGroupUsersLink(groupUserLinkCreateParams)
	if err != nil {
		return nil, err
	}

	groupUserLinkId, err := resInsertGroupUserLink.LastInsertId()
	if err != nil {
		return nil, err
	}

	groupsDependencies := GroupsDependencies{
		User:            &resInsertUser,
		Group:           &resInsertGroupUser,
		GroupUserLink:   &resInsertGroupUserLink,
		UserId:          userId,
		GroupUserId:     groupUserId,
		GroupUserLinkId: groupUserLinkId,
	}

	return &groupsDependencies, nil
}

func DeleteGroupDepencies(groupDependencies *GroupsDependencies) error {
	_, err := users.DeleteUser(int32(groupDependencies.UserId))
	if err != nil {
		return err
	}

	_, err = DeleteGroupsUserLink(int32(groupDependencies.GroupUserLinkId))
	if err != nil {
		return err
	}

	_, err = aadgu.DeleteGroupUser(int32(groupDependencies.GroupUserId))
	if err != nil {
		return err
	}

	return nil
}
