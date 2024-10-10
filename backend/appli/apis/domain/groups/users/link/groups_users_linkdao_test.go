package link

import (
	"database/sql"
	"fmt"
	"testing"

	aadgu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/groups/users"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestGroupsUserLinkDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	groupsDependencies, err := GetGroupsDependencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserLink, err := GetGroupUserLink(int32(groupsDependencies.GroupUserLinkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, groupUserLink.GulID, int32(groupsDependencies.GroupUserLinkId))

	CreateGroupUserParams := adm.CreateGroupUserParams{
		GuName:        "OTHER_TEST_STATUS",
		GuDescription: adm.NullString{NullString: sql.NullString{String: "Test Groups Status", Valid: true}},
	}

	res, err := aadgu.InsertGroupUser(CreateGroupUserParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	newGroupUserID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserLinkByGroupNameParams := adm.CreateGroupUserLinkByGroupNameParams{
		GulUserID: int32(groupsDependencies.UserId),
		GuName:    "OTHER_TEST_STATUS",
	}

	res, err = InsertGroupUserLinkByGroupName(groupUserLinkByGroupNameParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserLinkByGroupNameId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserLinkByGroupeName, err := GetGroupUserLink(int32(groupUserLinkByGroupNameId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupUserLinkByGroupeName.GulID, int32(groupUserLinkByGroupNameId))

	groupUserLinkUpdateParams := adm.UpdateGroupUserLinkParams{
		GulUserID:  int32(groupsDependencies.UserId),
		GulGroupID: int32(groupsDependencies.GroupUserId),
	}

	_, err = UpdateGroupUsersLink(groupUserLinkUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupsUserLinks, err := ListGroupsUsersLink()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupsUserLinks[0].GulGroupID, int32(groupsDependencies.GroupUserId))

	groupListByUser, err := aadgu.ListGroupsUsersByUser(int32(groupsDependencies.UserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, groupListByUser[0].GuDescription.String, "Test Groups Status")

	rows, err := DeleteGroupsUserLink(int32(groupUserLinkByGroupNameId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

	rows, err = aadgu.DeleteGroupUser(int32(newGroupUserID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

	err = DeleteGroupDepencies(groupsDependencies)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
