package datatypes

import (
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

type Privileges struct {
	AppliPrivileges []adm.GroupsPrivilege
}

type UserLogin struct {
	AppliUserLogin *aadu.AppliUserLogin
}

type PrivilegeParams struct {
	C             *gin.Context
	RequestMethod string
	Privilege     string
	HasPrivilege  bool
	Ul            UserLogin
	ParamId       string
	Check         func(UserLogin, int32) (bool, error)
}
