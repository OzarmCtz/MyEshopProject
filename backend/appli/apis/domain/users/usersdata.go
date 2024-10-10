package users

import (
	ma "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

type AppliUserLogin struct {
	User       ma.User              `json:"user"`
	LoginTime  string               `json:"login_time"`
	Privileges []ma.GroupsPrivilege `json:"privileges"`
	Groups     []ma.GroupsUser      `json:"groups_users"`
}
