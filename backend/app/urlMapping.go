package app

import (
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	"github.com/OzarmCtz/e_shop_backend_v1/app/middlewares"
	aacas "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/app/settings"
	aacda "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/dashboard"
	aacd "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/discount"
	aacdl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/discount/link"
	aagu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/groups/users"
	aagul "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/groups/users/link"
	aaci "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items"
	aacic "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items/category"
	aacicl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items/category/link"
	aacir "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items/reviews"
	aacisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items/sub/category"
	aaciscl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/items/sub/category/link"
	aacu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/users"
	aacub "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/users/basket"
	aacuw "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers/users/wishlist"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var middlewareList = []gin.HandlerFunc{
	middlewares.MainMiddleware(),
	middlewares.RouterMiddleware(),
	middlewares.PrivilegesMiddleware(),
}

func mapAppliUsersUrls() {

	// PRIVATE PATHS

	usersPrivateRoutes := router.Group(adg.PRIVATE_PATH + "/users")
	usersPrivateRoutes.GET("", aacu.List)
	usersPrivateRoutes.GET("/:"+adg.USER_ID, aacu.Get)
	usersPrivateRoutes.DELETE("/:"+adg.USER_ID, aacu.Delete)
	usersPrivateRoutes.PUT("/:"+adg.USER_ID, aacu.Update)

	appSettingsPrivatesRoutes := router.Group(adg.PRIVATE_PATH + "/app/settings")
	appSettingsPrivatesRoutes.GET("/:"+adg.APP_SETTINGS_ID, aacas.GetAppSettingsPrivate)
	appSettingsPrivatesRoutes.GET("", aacas.ListAppSettingPrivate)
	appSettingsPrivatesRoutes.PUT("/:"+adg.APP_SETTINGS_ID, aacas.UpdateAppSettingPrivate)
	appSettingsPrivatesRoutes.POST("", aacas.InsertAppSettingsPrivate)
	appSettingsPrivatesRoutes.DELETE("/:"+adg.APP_SETTINGS_ID, aacas.DeleteAppSettingPrivate)

	//--------------------------------------------\\

	usersPrivateUsersWishListRoutes := usersPrivateRoutes.Group("/wishlist")
	usersPrivateUsersWishListRoutes.GET("/:"+adg.USER_WISHLIST_ID, aacuw.GetItemsInUserWishListPrivate)
	usersPrivateUsersWishListRoutes.POST("", aacuw.InserUserWishListPrivate)
	usersPrivateUsersWishListRoutes.DELETE("", aacuw.DeleteItemPrivate)

	usersPrivateBasketRoutes := usersPrivateRoutes.Group("/basket")
	usersPrivateBasketRoutes.GET("/:"+adg.USER_BASKET_ID, aacub.GetItemsInUserWishListPrivate)
	usersPrivateBasketRoutes.POST("", aacub.InserUserBasketPrivate)
	usersPrivateBasketRoutes.DELETE("", aacub.DeleteUserBasketPrivate)

	groupsPrivateUsersLinkRoutes := router.Group(adg.PRIVATE_PATH + "/groups/users/link")
	groupsPrivateUsersLinkRoutes.GET("/:"+adg.GROUP_USER_LINK_ID, aagul.Get)
	groupsPrivateUsersLinkRoutes.PUT("/:"+adg.GROUP_USER_LINK_ID, aagul.Update)

	groupsPrivateUsersRoutes := router.Group(adg.PRIVATE_PATH + "/groups/users")
	groupsPrivateUsersRoutes.GET("", aagu.ListGroupUserPrivate)

	//--------------------------------------------\\

	discountsPrivateRoutes := router.Group(adg.PRIVATE_PATH + "/discount")
	discountsPrivateRoutes.GET("", aacd.ListDiscountPrivate)
	discountsPrivateRoutes.GET("/:"+adg.DISCOUNT_ID, aacd.GetDiscountPrivate)
	discountsPrivateRoutes.PUT("/:"+adg.DISCOUNT_ID, aacd.UpdateDiscountPrivate)
	discountsPrivateRoutes.POST("", aacd.InsertDiscountPrivate)
	discountsPrivateRoutes.DELETE("/:"+adg.DISCOUNT_ID, aacd.DeleteDiscountPrivate)

	discountsPrivateLinkRoutes := discountsPrivateRoutes.Group("/link")
	discountsPrivateLinkRoutes.GET("", aacdl.ListDiscountLinkPrivate)
	discountsPrivateLinkRoutes.GET("/:"+adg.DISCOUNT_LINK_ID, aacdl.GetDiscountLinkPrivate)
	discountsPrivateLinkRoutes.PUT("/:"+adg.DISCOUNT_LINK_ID, aacdl.UpdateDiscountLinkPrivate)
	discountsPrivateLinkRoutes.POST("", aacdl.InsertDiscountLinkPrivate)
	discountsPrivateLinkRoutes.DELETE("/:"+adg.DISCOUNT_LINK_ID, aacdl.DeleteDiscountLinkPrivate)

	itemsPrivateRoutes := router.Group(adg.PRIVATE_PATH + "/items")
	itemsPrivateRoutes.GET("", aaci.ListItemPrivate)
	itemsPrivateRoutes.GET("/:"+adg.ITEM_ID, aaci.GetItemPrivate)
	itemsPrivateRoutes.PUT("/:"+adg.ITEM_ID, aaci.UpdateItemPrivate)
	itemsPrivateRoutes.POST("", aaci.InsertItemPrivate)
	itemsPrivateRoutes.DELETE("/:"+adg.ITEM_ID, aaci.DeleteItemPrivate)

	reviewsPrivateRoutes := itemsPrivateRoutes.Group("/reviews")
	reviewsPrivateRoutes.GET("/:"+adg.ITEM_REVIEW_ID, aacir.ListItemReviewByUserIdPrivate)
	reviewsPrivateRoutes.GET("", aacir.ListItemReviewPrivate)
	reviewsPrivateRoutes.POST("", aacir.InsertItemReviewPrivate)
	reviewsPrivateRoutes.DELETE("/:"+adg.ITEM_REVIEW_ID, aacir.DeleteItemeReviewPrivate)

	itemsCategoryPrivateRoutes := itemsPrivateRoutes.Group("/category")
	itemsCategoryPrivateRoutes.POST("", aacic.InsertItemCategoryPrivate)
	itemsCategoryPrivateRoutes.PUT("/:"+adg.ITEM_CATEGORY_ID, aacic.UpdateItemCategoryPrivate)
	itemsCategoryPrivateRoutes.DELETE("/:"+adg.ITEM_CATEGORY_ID, aacic.DeleteItemCategoryPrivate)
	itemsCategoryPrivateRoutes.GET("", aacic.ListItemSubCategoryPrivate)

	itemsCategoryLinkPrivateRoutes := itemsCategoryPrivateRoutes.Group("/link")
	itemsCategoryLinkPrivateRoutes.GET("/:"+adg.ITEM_CATEGORY_LINK_ID, aacicl.GetItemCategoryLinkPrivate)
	itemsCategoryLinkPrivateRoutes.GET("", aacicl.ListItemCategoryLinkPrivate)
	itemsCategoryLinkPrivateRoutes.PUT("/:"+adg.ITEM_CATEGORY_LINK_ID, aacicl.UpdateItemCategoryLinkPrivate)
	itemsCategoryLinkPrivateRoutes.POST("", aacicl.InsertItemCategoryLinkPrivate)
	// itemsCategoryLinkPrivateRoutes.DELETE("/:"+adg.ITEM_CATEGORY_LINK_ID, aacicl.DeleteItemCategoryLinkPrivate)

	itemsSubCategoryPrivateRoutes := itemsPrivateRoutes.Group("/sub/category")
	itemsSubCategoryPrivateRoutes.POST("", aacisc.InsertSubItemCategoryPrivate)
	itemsSubCategoryPrivateRoutes.PUT("/:"+adg.ITEM_SUB_CATEGORY_ID, aacisc.UpdateItemSubCategoryPrivate)
	itemsSubCategoryPrivateRoutes.DELETE("/:"+adg.ITEM_SUB_CATEGORY_ID, aacisc.DeleteItemSubCategoryPrivate)
	itemsSubCategoryPrivateRoutes.GET("/:"+adg.ITEM_SUB_CATEGORY_ID, aacisc.GetItemSubCategoryPrivate)
	itemsSubCategoryPrivateRoutes.GET("", aacisc.ListItemSubCategoryPrivate)

	itemsSubCategoryLinkPrivatesRoutes := itemsSubCategoryPrivateRoutes.Group("/link")
	itemsSubCategoryLinkPrivatesRoutes.GET("/:"+adg.ITEM_SUB_CATEGORY_LINK_ID, aaciscl.GetItemSubCategoryLinkPrivate)
	itemsSubCategoryLinkPrivatesRoutes.GET("", aaciscl.ListItemSubCategoryLinkPrivate)
	itemsSubCategoryLinkPrivatesRoutes.POST("", aaciscl.InsertItemSubCategoryLinkPrivate)
	itemsSubCategoryLinkPrivatesRoutes.PUT("/:"+adg.ITEM_SUB_CATEGORY_LINK_ID, aaciscl.UpdateItemCategoryLinkPrivate)
	itemsSubCategoryLinkPrivatesRoutes.DELETE("/:"+adg.ITEM_SUB_CATEGORY_LINK_ID, aaciscl.UpdateItemCategoryLinkPrivate)

	dashboardPublicRoutes := router.Group(adg.PRIVATE_PATH + "/dashboard")
	dashboardPublicRoutes.GET("", aacda.GetAccess)

	// PUBLIC PATHS
	signUpRouter := router.Group(adg.PUBLIC_PATH + "/sign/up")
	signUpRouter.POST("", aacu.SignUpUser)

	//--------------------------------------------\\

	appSettingsPublicRoutes := router.Group(adg.PUBLIC_PATH + "/app/settings")
	appSettingsPublicRoutes.GET("", aacas.ListAppSettingsPublic)
	appSettingsPublicRoutes.GET("/:appKey", aacas.GetAppSettingsByKeyPublic)

	itemsPublicRoutes := router.Group(adg.PUBLIC_PATH + "/items")
	itemsPublicRoutes.GET("", aaci.ListItemPublic)
	itemsPublicRoutes.GET("/:"+adg.ITEM_ID, aaci.GetItemPublic)

	reviewsPublicRoutes := itemsPublicRoutes.Group("/reviews")
	reviewsPublicRoutes.GET("/:"+adg.ITEM_REVIEW_ID, aacir.ListItemReviewByItemIdPublic)

	itemsSubCategoryPublicRoutes := itemsPublicRoutes.Group("/sub/category")
	itemsSubCategoryPublicRoutes.GET("/:"+adg.ITEM_SUB_CATEGORY_ID, aacisc.GetItemSubCategoryPublic) // useless ?
	itemsSubCategoryPublicRoutes.GET("", aacisc.ListItemSubCategoryPublic)                           // useless ?

	itemsSubCategoryLinkPublicRoutes := itemsSubCategoryPublicRoutes.Group("/link")
	itemsSubCategoryLinkPublicRoutes.GET("/:"+adg.ITEM_SUB_CATEGORY_LINK_ID, aaciscl.ListItemsBySubCategoryPublic)

	itemsCategoryPublicRoutes := itemsPublicRoutes.Group("/category")
	itemsCategoryPublicRoutes.GET("", aacic.ListItemCategoryPublic)
	itemsCategoryPublicRoutes.GET("/:"+adg.ITEM_CATEGORY_ID, aacic.GetItemCategoryPublic) // useless ?

}

func MapUrls() {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "route not found")
	})

	router.SetTrustedProxies(nil)

	gin.SetMode(gin.ReleaseMode)

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(corsConfig))

	router.Use(middlewareList...)
	mapAppliUsersUrls()
}
