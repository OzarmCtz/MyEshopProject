package privileges

import (
	"database/sql"
	"fmt"
	"testing"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
)

func TestGroupsPrivilegesDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	resInsertPrivilege, err := InsertGroupPrivilege(adg.NO_CLIENT_QUERIES_DB_ID, adm.NullString{NullString: sql.NullString{String: "Privilege", Valid: true}})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	privId, err := resInsertPrivilege.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	privilege, err := GetGroupPrivilege(adg.NO_CLIENT_QUERIES_DB_ID, int32(privId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, privilege.GpID, int32(privId))

	groupUpdateParams := adm.UpdateGroupPrivilegeParams{
		GpID:   int32(privId),
		GpPath: adm.NullString{NullString: sql.NullString{String: "Privilege updated", Valid: true}},
	}

	rows, err := UpdateGroupPrivilege(adg.NO_CLIENT_QUERIES_DB_ID, groupUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

	privileges, err := ListGroupsPrivileges(adg.NO_CLIENT_QUERIES_DB_ID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println((privileges))

	for _, privilege := range privileges {
		if privilege.GpID == int32(privId) {
			assert.Equal(t, privilege.GpPath, adm.NullString{NullString: sql.NullString{String: "Privilege updated", Valid: true}})
		}
	}

	rows, err = DeleteGroupPrivilege(adg.NO_CLIENT_QUERIES_DB_ID, privilege.GpID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, rows, int64(1))

}
