package users

import (
	"fmt"
	"testing"
	"time"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestUserDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	// Create User
	user, err := CreateUser()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, user.UID, int32(user.UID))

	// UpdateUser
	userUpdateParams := adm.UpdateUserParams{
		UEmail:        "johndoev2@gmail.com",
		URegisterDate: time.Now(),
		UIsDisabled:   false,
		UID:           int32(user.UID),
	}

	rows, err := UpdateUser(userUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	// 	Get Updated User
	users, err := ListUsers()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, users[0].UEmail, "johndoev2@gmail.com")

	// Get User By Email
	userByEmail, err := GetUserByEmail("johndoev2@gmail.com")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, userByEmail.UID, int32(1))

	// Attempt SQL Injection
	injectionEmail := "invalid' OR '1'='1"
	_, err = GetUserByEmail(injectionEmail)
	if err == nil {
		fmt.Println("SQL Injection should not succeed")
		t.Fail()
	}

}
