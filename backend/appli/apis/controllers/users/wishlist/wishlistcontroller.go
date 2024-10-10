package wishlist

import (
	"fmt"
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasuw "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/users/wishlist"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func GetItemsInUserWishListPrivate(c *gin.Context) {

	wishListUsersId, wishListUserIdExist := c.Get(adg.USER_WISHLIST_ID)
	userLogin, userExist := c.Get("userLogin")

	if userExist && wishListUserIdExist {
		wishListUserId := wishListUsersId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)

		items, err := aasuw.WishListService.GetUserWishListByUserPrivate(currentUser, wishListUserId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(items) == 0 {
			c.JSON(http.StatusNotFound, items)
			return
		}

		c.JSON(http.StatusOK, items)
		return
	}

	if !userExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !wishListUserIdExist {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func InserUserWishListPrivate(c *gin.Context) {
	userLogin, userLoginExist := c.Get("userLogin")
	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		userWishListParams := adm.CreateUserWishListParams{}
		if err := c.ShouldBindJSON(&userWishListParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res, err := aasuw.WishListService.InsertUserWishListPrivate(currentUser, userWishListParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func DeleteItemPrivate(c *gin.Context) {
	userLogin, userExists := c.Get("userLogin")

	if userExists {
		currentUser := userLogin.(aadu.AppliUserLogin)

		userWishListDeleteParams := adm.DeleteUserWishListByUserAndItemsParams{}
		if err := c.ShouldBindJSON(&userWishListDeleteParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasuw.WishListService.DeleteUserWishListByUserAndItemPrivate(currentUser, userWishListDeleteParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		isErr := aac.EditErrorResponse(rows, userWishListDeleteParams.WlUserID, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item wish list %d deleted ", userWishListDeleteParams.WlUserID))
		}
		return
	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}
