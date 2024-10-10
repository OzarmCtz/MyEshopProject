package privileges

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
)

func GetGroupPrivilegesByUserId(userId int32) ([]adm.GroupsPrivilege, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	privileges, err := adm.QueriesDb.GetGroupPrivilegesByUserId(ctx, userId)
	return privileges, err
}

func ListGroupsPrivileges(userClientId int32) ([]adm.GroupsPrivilege, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	privileges, err := aum.GetQueriesDbFromDbPool(userClientId).ListGroupsPrivileges(ctx)
	return privileges, err
}

func GetGroupPrivilege(userClientId int32, privId int32) (adm.GroupsPrivilege, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	privilege, err := aum.GetQueriesDbFromDbPool(userClientId).GetGroupPrivilege(ctx, privId)
	return privilege, err
}

func InsertGroupPrivilege(userClientId int32, privilege adm.NullString) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := aum.GetQueriesDbFromDbPool(userClientId).CreateGroupPrivilege(ctx, privilege)
	return res, err
}

func UpdateGroupPrivilege(userClientId int32, privilegeParams adm.UpdateGroupPrivilegeParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := aum.GetQueriesDbFromDbPool(userClientId).UpdateGroupPrivilege(ctx, privilegeParams)
	return rows, err
}

func DeleteGroupPrivilege(userClientId int32, privId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := aum.GetQueriesDbFromDbPool(userClientId).DeleteGroupPrivilege(ctx, privId)
	return rows, err
}
