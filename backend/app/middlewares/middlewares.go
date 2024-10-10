package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adt "github.com/OzarmCtz/e_shop_backend_v1/app/data/types"
	aug "github.com/OzarmCtz/e_shop_backend_v1/app/utils/gin"
	"github.com/OzarmCtz/e_shop_backend_v1/app/utils/logger"
	aur "github.com/OzarmCtz/e_shop_backend_v1/app/utils/resterrors"
	aurl "github.com/OzarmCtz/e_shop_backend_v1/app/utils/url"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	"github.com/gin-gonic/gin"
)

var statusUnauthorizedTriggered = false
var startTime time.Time

func MainMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime = time.Now()
		c.Next()
	}
}

func LogsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if isAppliPrivatePath(c.Request.URL.Path) {
			userLogin, userLoginExists := c.Get("userLogin")

			end := time.Now()

			token := c.Query("token")

			if userLoginExists {
				currentUser := userLogin.(aadu.AppliUserLogin)
				user, err := aug.GetAplliUserFromContext(currentUser, token)
				if err != nil {
					if !statusUnauthorizedTriggered {
						c.JSON(http.StatusUnauthorized, adg.UNAUTHORIZED)
					}
					c.Abort()
					return
				}
				latency := end.Sub(startTime)
				entry := logger.RequestLog{
					Method:  c.Request.Method,
					Path:    c.Request.URL.Path,
					User:    strconv.FormatInt(int64(user.UID), 10),
					Latency: fmt.Sprintf("%v", latency),
				}
				logger.RecordAppliEvent(currentUser, entry)
			}

		}

		c.Next()
	}
}

// RouterMiddleware is a middleware function that handles routing logic.
func RouterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if isAppliPublicPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		if isAppliPrivatePath(c.Request.URL.Path) {
			token := c.Query("token")
			resterr := AppliAuth(c, token)
			if resterr != nil {
				c.JSON(resterr.Status(), resterr.Error())
				c.Abort()
				return
			}
			var fakeappliUserLogin aadu.AppliUserLogin

			user, err := aug.GetAplliUserFromContext(fakeappliUserLogin, token)

			if err != nil {
				statusUnauthorizedTriggered = true
				c.JSON(http.StatusUnauthorized, "appli user: "+err.Error())
				c.Abort()
				return
			}

			err = aug.SaveAppliUserLoginFromContext(c, user)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
				return
			}
			c.Next()
			return
		}

		err := errors.New("bad request path")
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()

	}

}

func AppliAuth(c *gin.Context, token string) aur.RestError {

	_, err := aug.GetCentralUserFromContext(token)
	if err != nil {
		resterr := aur.NewUnauthorizedError(adg.AUTH_ERROR, err)
		return resterr
	}

	return nil

}

func PrivilegesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err aur.RestError

		resourceName := ""
		if isAppliPrivatePath(c.Request.URL.Path) {
			resourceName = aurl.GetResourceName(c.Request.URL.Path, adg.APPLI_PRIVATE_PATH_SUBMATCH_REGEX)
			privilege := GetPrivilegeByMethod(c.Request.Method, resourceName)
			err = checkPrivileges(c, resourceName, privilege)
			if err != nil {
				c.JSON(err.Status(), err.Error())
				c.Abort()
				return
			}

		} else if isAppliPublicPath(c.Request.URL.Path) {
			return
		} else {
			err := aur.NewNotFoundError(adg.PATH_NOT_FOUND_ERROR, errors.New(adg.PATH_NOT_FOUND_MESSAGE))
			c.JSON(err.Status(), err.Error())
			c.Abort()
			return
		}

		c.Next()

	}
}

func checkPrivileges(c *gin.Context, resourceName, privilege string) aur.RestError {
	var aul aadu.AppliUserLogin
	var ul adt.UserLogin
	var err aur.RestError

	userLogin, _ := c.Get(adg.USER_LOGIN)

	if isAppliPrivatePath(c.Request.URL.Path) {
		aul = userLogin.(aadu.AppliUserLogin)
		c.Set("userLogin", aul)
		ul = adt.UserLogin{AppliUserLogin: &aul}

	} else {
		return aur.NewNotFoundError(adg.PATH_NOT_FOUND_ERROR, errors.New(adg.PATH_NOT_FOUND_MESSAGE))
	}

	switch resourceName {
	case adg.USERS:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.USER_ID, nil)
	case adg.GROUPS_USERS:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.GROUP_USER_ID, nil)
	case adg.GROUPS_USERS_LINK:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.GROUP_USER_LINK_ID, nil)
	case adg.ITEMS:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_ID, nil)
	case adg.ITEMS_CATEGORY:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_CATEGORY_ID, nil)
	case adg.ITEMS_SUB_CATEGORY:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_SUB_CATEGORY_ID, nil)
	case adg.ITEMS_SUB_CATEGORY_LINK:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_SUB_CATEGORY_LINK_ID, nil)
	case adg.ITEMS_CATEGORY_LINK:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_CATEGORY_LINK_ID, nil)
	case adg.USERS_WISHLIST:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.USERS_WISHLIST_ID, nil)
	case adg.USERS_BASKET:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.USER_BASKET_ID, nil)
	case adg.ITEMS_REVIEWS:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.ITEM_REVIEW_ID, nil)
	case adg.APP_SETTINGS:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.APP_SETTINGS_ID, nil)
	case adg.DISCOUNT:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.DISCOUNT_ID, nil)
	case adg.DISCOUNT_LINK:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.DISCOUNT_LINK_ID, nil)
	case adg.DASHBOARD:
		err = checkPrivilegesByResourceName(c, ul, privilege, adg.DASHBOARD_ID, nil)

	default:
		err = aur.NewNotFoundError(adg.PRIVILEGE_NOT_FOUND_MESSAGE, errors.New(adg.PRIVILEGE_NOT_FOUND_ERROR))

	}

	if err != nil {
		return err
	}

	return nil

}

func isAppliPrivatePath(path string) bool {
	re := regexp.MustCompile(adg.APPLI_PRIVATE_PATH_REGEX)
	return re.Match([]byte(path))
}

func isAppliPublicPath(path string) bool {
	re := regexp.MustCompile(adg.APPLI_PUBLIC_PATH_REGEX)
	return re.Match([]byte(path))
}
