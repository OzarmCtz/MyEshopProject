package global

import "time"

var (
	TIMEOUT                           = 2000
	Timeout                           = time.Duration(TIMEOUT) * time.Second
	NO_CLIENT_QUERIES_DB_ID           = int32(0)
	NO_CLIENT_ID                      = int32(0)
	UNAUTHORIZED                      = "unauthorized"
	SUPER_ADMIN_STATUS                = "SUPERADMIN_STATUS"
	ADMIN_STATUS                      = "ADMIN_STATUS"
	AUTH_ERROR                        = "authentication error"
	APPLI_PRIVATE_PATH_SUBMATCH_REGEX = `\/api\/v1\/private\/(.*)$`
	APPLI_PUBLIC_PATH_SUBMATCH_REGEX  = `\/api\/v1\/public\/(.*)$`
	APPLI_PRIVATE_PATH_REGEX          = `\/api\/v1\/private\/.*$`
	APPLI_PUBLIC_PATH_REGEX           = `\/api\/v1\/public\/.*$`
	PATH_NOT_FOUND_ERROR              = "Path not found"
	PATH_NOT_FOUND_MESSAGE            = "the path does not exist for this resource"
	GROUPS_USERS                      = "groups/users"
	GROUPS_USERS_LINK                 = "groups/users/link"
	ITEMS                             = "items"
	ITEMS_CATEGORY                    = "items/category"
	ITEMS_SUB_CATEGORY                = "items/sub/category"
	ITEMS_CATEGORY_LINK               = "items/category/link"
	ITEMS_SUB_CATEGORY_LINK           = "items/sub/category/link"
	USERS_WISHLIST                    = "users/wishlist"
	USERS_BASKET                      = "users/basket"
	ITEMS_REVIEWS                     = "items/reviews"
	APP_SETTINGS                      = "app/settings"
	DISCOUNT                          = "discount"
	GET                               = "GET"
	POST                              = "POST"
	READ                              = "READ"
	CREATE                            = "CREATE"
	PATCH                             = "PATCH"
	PUT                               = "PUT"
	DELETE                            = "DELETE"
	UPDATE                            = "UPDATE"
	USER_LOGIN                        = "userLogin"
	USER_ERROR_MESSAGE                = "User related to the request was not found"
	PRIVATE_PATH                      = "/api/v1/private"
	PUBLIC_PATH                       = "/api/v1/public"
	USERS                             = "users"
	DASHBOARD                         = "dashboard"
	USER_ID                           = "user_id"
	USER_WISHLIST_ID                  = "users_wishlist_id"
	APP_SETTINGS_ID                   = "app_settings_id"
	USER_BASKET_ID                    = "users_basket_id"
	ITEM_ID                           = "item_id"
	DISCOUNT_ID                       = "discount_id"
	DISCOUNT_LINK_ID                  = "discount_link_id"
	DASHBOARD_ID                      = "dashboard_id"
	DISCOUNT_LINK                     = "discount/link"
	ITEM_REVIEW_ID                    = "items_reviews_id"
	ITEM_CATEGORY_ID                  = "item_category_id"
	ITEM_SUB_CATEGORY_LINK_ID         = "item_sub_category_link_id"
	USERS_WISHLIST_ID                 = "users_wishlist_id"
	ITEM_CATEGORY_LINK_ID             = "item_category_link_id"
	ITEM_SUB_CATEGORY_ID              = "item_sub_category_id"
	GROUP_USER_LINK_ID                = "group_user_link_id"
	GROUP_USER_ID                     = "group_user_id"
	PRIVILEGE_NOT_FOUND_ERROR         = "checking privileges for the resource cannot not be proceed "
	PRIVILEGE_NOT_FOUND_MESSAGE       = "Can not find privileges for the resource name "
	PARAM_ERROR_MESSAGE               = "Are you sure the correct parameter was passed to the path: %v"
	NO_ROWS_IN_RESULT_SET_ERROR       = "sql: no rows in result set"
	CLIENT_STATUS                     = "CLIENT_STATUS"
)
