package link

import (
	"fmt"
	"net/http"

	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"

	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasdl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/discount/link"

	"github.com/gin-gonic/gin"
)

func GetDiscountLinkPrivate(c *gin.Context) {

	discountLinksId, discountLinkIdExists := c.Get(adg.DISCOUNT_LINK_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if discountLinkIdExists && userLoginExist {
		discountLinkId := discountLinksId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		item, err := aasdl.DiscountLinkService.GetDiscountLinkPrivate(currentUser, discountLinkId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, item)

	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !discountLinkIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func ListDiscountLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		discounts, err := aasdl.DiscountLinkService.ListDiscountsLinksPrivate(currentUser)

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

func InsertDiscountLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		var discountLinkParams adm.CreateDiscountLinkParams
		err := c.ShouldBindJSON(&discountLinkParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		discountLink, err := aasdl.DiscountLinkService.InsertDiscountLinkPrivate(currentUser, discountLinkParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, discountLink)
		return
	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func UpdateDiscountLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")
	discountLinkId, discountIdExist := c.Get(adg.DISCOUNT_LINK_ID)

	if userExist && discountIdExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		dlId := discountLinkId.(int32)
		updateMysqlDiscountLink := adm.UpdateDiscountLinkParams{DlID: dlId}

		if err := c.ShouldBindJSON(&updateMysqlDiscountLink); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasdl.DiscountLinkService.UpdateDiscountLinkPrivate(currentUser, updateMysqlDiscountLink)
		isErr := aac.EditErrorResponse(rows, dlId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("discount link %d updated ", dlId))
		}
		return
	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !discountIdExist {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func DeleteDiscountLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")
	discountLinkId, discountLinkIdExist := c.Get(adg.DISCOUNT_LINK_ID)

	if userExist && discountLinkIdExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		dlId := discountLinkId.(int32)

		rows, err := aasdl.DiscountLinkService.DeleteDiscountLinkPrivate(currentUser, dlId)
		isErr := aac.EditErrorResponse(rows, dlId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("discount link %d deleted ", dlId))
		}

		return

	}

	if !userExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !discountLinkIdExist {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}
