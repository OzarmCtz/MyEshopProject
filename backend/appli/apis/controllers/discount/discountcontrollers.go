package discount

import (
	"fmt"
	"net/http"

	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasd "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/discount"

	"github.com/gin-gonic/gin"
)

func GetDiscountPrivate(c *gin.Context) {

	discountsId, discountIdExists := c.Get(adg.DISCOUNT_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if discountIdExists && userLoginExist {
		discountId := discountsId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		item, err := aasd.DiscountService.GetDiscountPrivate(currentUser, discountId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, item)

	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !discountIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func ListDiscountPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		discounts, err := aasd.DiscountService.ListDiscountsPrivate(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(discounts) == 0 {
			c.JSON(http.StatusNotFound, discounts)
			return
		}

		c.JSON(http.StatusOK, discounts)
		return
	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func InsertDiscountPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")
	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		discountParams := adm.CreateDiscountParams{}
		if err := c.ShouldBindJSON(&discountParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res, err := aasd.DiscountService.InsertDiscountPrivate(currentUser, discountParams)
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

func UpdateDiscountPrivate(c *gin.Context) {
	discountsId, discountsIdexists := c.Get(adg.DISCOUNT_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if discountsIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		dId := discountsId.(int32)
		updateMysqlDiscount := adm.UpdateDiscountParams{DID: dId}
		if err := c.ShouldBindJSON(&updateMysqlDiscount); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasd.DiscountService.UpdateDiscountPrivate(currentUser, updateMysqlDiscount)
		isErr := aac.EditErrorResponse(rows, dId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("discount %d updated ", dId))
		}
		return
	}
	if !discountsIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func DeleteDiscountPrivate(c *gin.Context) {
	discountsId, discountIdExists := c.Get(adg.DISCOUNT_ID)
	userLogin, userExists := c.Get("userLogin")

	if discountIdExists && userExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		dId := discountsId.(int32)
		rows, err := aasd.DiscountService.DeleteDiscountPrivate(currentUser, dId)
		isErr := aac.EditErrorResponse(rows, dId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("discount %d deleted ", dId))
		}
		return
	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !discountIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}
