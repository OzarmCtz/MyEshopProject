package middlewares

import (
	"errors"
	"fmt"
	"strings"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adt "github.com/OzarmCtz/e_shop_backend_v1/app/data/types"
	aug "github.com/OzarmCtz/e_shop_backend_v1/app/utils/gin"
	aur "github.com/OzarmCtz/e_shop_backend_v1/app/utils/resterrors"
	aua "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth"
	"github.com/gin-gonic/gin"
)

func GetPrivilegeByMethod(method, resourceName string) string {
	if strings.Contains(resourceName, adg.GROUPS_USERS) {
		resourceNameSplit := strings.Split(resourceName, "_")
		resourceName = resourceNameSplit[0]
	}
	privilege := resourceName + "_"

	if method == adg.GET {
		privilege += adg.READ
	}
	if method == adg.POST {
		privilege += adg.CREATE
	}
	if method == adg.PUT || method == adg.PATCH {
		privilege += adg.UPDATE
	}
	if method == adg.DELETE {
		privilege += adg.DELETE
	}
	return strings.ToUpper(privilege)
}

func checkPrivilegesByResourceName(c *gin.Context, ul adt.UserLogin, privilege, paramId string,
	check func(userLogin adt.UserLogin, paramClientId int32) (bool, error)) aur.RestError {
	privilegeParams := &adt.PrivilegeParams{
		C:             c,
		RequestMethod: c.Request.Method,
		Privilege:     privilege,
		HasPrivilege:  false,
		Ul:            ul,
		ParamId:       paramId,
		Check:         check,
	}
	return CheckPrivileges(c, privilegeParams)
}

func requestMethodMatchMethod(hasPrivilegeParams *adt.PrivilegeParams, method string) bool {
	return hasPrivilegeParams.C.Request.Method == method
}

func CheckPrivileges(c *gin.Context, privilegeParams *adt.PrivilegeParams) aur.RestError {
	var err error

	userId := 0

	privilegeParams, err = EnsurehasPrivilege(privilegeParams)

	if requestMethodMatchMethod(privilegeParams, adg.POST) {
		privilegeParams.RequestMethod = adg.POST
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}

	if requestMethodMatchMethod(privilegeParams, adg.PUT) {
		privilegeParams.RequestMethod = adg.PUT
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}

	if requestMethodMatchMethod(privilegeParams, adg.DELETE) {
		privilegeParams.RequestMethod = adg.DELETE
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}

	if privilegeParams.Ul.AppliUserLogin != nil {
		userId = int(privilegeParams.Ul.AppliUserLogin.User.UID)
	}

	if err != nil {
		return aur.NewBadRequestError(fmt.Sprintf("error with %v method in route %v for user %d", privilegeParams.C.Request.Method, privilegeParams.C.Request.URL.Path, userId), err)
	}

	if !privilegeParams.HasPrivilege {
		//return aur.NewUnauthorizedError("error checking privileges", errors.New("user does not have privileges"))
		return aur.NewUnauthorizedError("", errors.New("user does not have privileges"))
	}
	return nil
}

func EnsurehasPrivilege(privilegesParams *adt.PrivilegeParams) (*adt.PrivilegeParams, error) {
	var privileges adt.Privileges
	hasFoundPrivileges := false

	if privilegesParams.Ul.AppliUserLogin != nil {
		privileges.AppliPrivileges = privilegesParams.Ul.AppliUserLogin.Privileges
		hasFoundPrivileges = aua.FindPrivilege(privileges.AppliPrivileges, privilegesParams.Privilege)
	}

	if hasFoundPrivileges {
		id, err := aug.GetParamId(privilegesParams.C, privilegesParams.ParamId)
		if err != nil {
			return privilegesParams, err
		}

		if id > 0 {
			if privilegesParams.Check != nil {
				privilegesParams.HasPrivilege, err = privilegesParams.Check(privilegesParams.Ul, id)
				if err != nil {
					return privilegesParams, err
				}
			} else {
				privilegesParams.HasPrivilege = true
			}
			privilegesParams.C.Set(privilegesParams.ParamId, id)
		} else {
			privilegesParams.HasPrivilege = true
		}
	}
	return privilegesParams, nil
}
