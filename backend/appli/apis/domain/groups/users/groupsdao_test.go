package users

import (
	"database/sql"
	"fmt"
	"testing"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestGroupsUsersDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	CreateGroupUserParams := adm.CreateGroupUserParams{
		GuName:        "TEST_STATUS",
		GuDescription: adm.NullString{NullString: sql.NullString{String: "Test Groups Status", Valid: true}},
	}

	resInsertGroupUser, err := InsertGroupUser(CreateGroupUserParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserId, err := resInsertGroupUser.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	group, err := GetGroupUser(int32(groupUserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, group.GuID, int32(groupUserId))

	groupUserByName, err := GetGroupUserByName("TEST_STATUS")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupUserByName.GuName, "TEST_STATUS")

	groupUsersUpdateParams := adm.UpdateGroupUserParams{
		GuName:        "TEST_STATUS_UPDATED",
		GuDescription: adm.NullString{NullString: sql.NullString{String: "Test Groups Status Updated", Valid: true}},
		GuID:          int32(groupUserId),
	}

	rows, err := UpdateGroupUser(groupUsersUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

	groupsUsers, err := ListGroupsUsers()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, groupsUsers[0].GuName, "TEST_STATUS_UPDATED")

	rows, err = DeleteGroupUser(int32(groupUserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

}
