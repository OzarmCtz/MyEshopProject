package basket

import (
	"fmt"
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasub "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/users/basket"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func GetItemsInUserWishListPrivate(c *gin.Context) {

	basketUsersId, basketUserIdExist := c.Get(adg.USER_BASKET_ID)
	userLogin, userExist := c.Get("userLogin")

	if userExist && basketUserIdExist {
		basketUserId := basketUsersId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)

		baskets, err := aasub.BasketService.GetUserBasketByUserPrivate(currentUser, basketUserId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(baskets) == 0 {
			c.JSON(http.StatusNotFound, baskets)
			return
		}

		c.JSON(http.StatusOK, baskets)
		return
	}

	if !userExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !basketUserIdExist {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func InserUserBasketPrivate(c *gin.Context) {
	userLogin, userLoginExist := c.Get("userLogin")
	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		userBasketParams := adm.CreateUserBasketParams{}
		if err := c.ShouldBindJSON(&userBasketParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res, err := aasub.BasketService.InsertUserBasketPrivate(currentUser, userBasketParams)
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

func DeleteUserBasketPrivate(c *gin.Context) {
	userLogin, userExists := c.Get("userLogin")

	if userExists {
		currentUser := userLogin.(aadu.AppliUserLogin)

		userBasketDeleteParams := adm.DeleteUserBasketByUserAndItemsParams{}
		if err := c.ShouldBindJSON(&userBasketDeleteParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasub.BasketService.DeleteUserBasketByUserAndItemPrivate(currentUser, userBasketDeleteParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		isErr := aac.EditErrorResponse(rows, userBasketDeleteParams.UbUserID, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item basket %d deleted ", userBasketDeleteParams.UbUserID))
		}
		return
	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}
