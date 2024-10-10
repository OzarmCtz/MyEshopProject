package categorycontroller

import (
	"fmt"
	"net/http"
	"strconv"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasic "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/category"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemSubCategoryPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		items, err := aasic.ItemsCategoryService.ListItemsCategoryPrivate(currentUser)

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

	c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

}

func ListItemCategoryPublic(c *gin.Context) {
	items, err := aasic.ItemsCategoryService.ListItemsCategoryPublic()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusNotFound, items)
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItemCategoryPublic(c *gin.Context) {

	itemCategoryIdExists := c.Param(adg.ITEM_CATEGORY_ID)

	if itemCategoryIdExists != "" {
		itemCategoryId, err := strconv.Atoi(itemCategoryIdExists)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item category ID format"})
			return
		}
		itemCategory, err := aasic.ItemsCategoryService.GetItemsCategoryPublic(int32(itemCategoryId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itemCategory)
		return
	}

	if itemCategoryIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func InsertItemCategoryPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")

	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		var itemCategoryParams adm.CreateItemCategoryParams
		if err := c.ShouldBindJSON(&itemCategoryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := aasic.ItemsCategoryService.InserItemsCategoryPrivate(currentUser, itemCategoryParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
		return
	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func UpdateItemCategoryPrivate(c *gin.Context) {

	itemsCategoryId, itemsCategoryIdexists := c.Get(adg.ITEM_CATEGORY_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsCategoryIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		icId := itemsCategoryId.(int32)

		updateMysqlItemCategory := adm.UpdateItemCategoryParams{IcID: icId}
		if err := c.ShouldBindJSON(&updateMysqlItemCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasic.ItemsCategoryService.UpdateItemsCategoryPrivate(currentUser, updateMysqlItemCategory)
		isErr := aac.EditErrorResponse(rows, icId, err, c)

		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item category %d updated ", icId))
		}
		return

	}
	if !itemsCategoryIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func DeleteItemCategoryPrivate(c *gin.Context) {

	userLogin, userExists := c.Get("userLogin")
	itemsCategoryId, itemCategoryIdExists := c.Get(adg.ITEM_CATEGORY_ID)

	if userExists && itemCategoryIdExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		itemCategoryId := itemsCategoryId.(int32)

		rows, err := aasic.ItemsCategoryService.DeleteItemsCategoryPrivate(currentUser, itemCategoryId)
		isErr := aac.EditErrorResponse(rows, itemCategoryId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item category %d deleted ", itemCategoryId))
		}
		return

	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemCategoryIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}
