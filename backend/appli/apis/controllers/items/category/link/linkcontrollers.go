package linkcontroller

import (
	"fmt"
	"net/http"
	"strconv"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasicl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/category/link"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemSubCategoryByCategoryPublic(c *gin.Context) {

	itemSubCategoryIdExists := c.Param(adg.ITEM_SUB_CATEGORY_ID)

	if itemSubCategoryIdExists != "" {
		itemSubCategoryId, err := strconv.Atoi(itemSubCategoryIdExists)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item category ID format"})
			return
		}

		items, err := aasicl.CategoryLinkService.ListItemsCategoryByCategoryPublic(int32(itemSubCategoryId))

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

	if itemSubCategoryIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func ListItemCategoryLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		itemsCategoryLink, err := aasicl.CategoryLinkService.ListItemsCategoryLinkPrivate(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(itemsCategoryLink) == 0 {
			c.JSON(http.StatusNotFound, itemsCategoryLink)
			return
		}

		c.JSON(http.StatusOK, itemsCategoryLink)
		return

	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

}

func GetItemCategoryLinkPrivate(c *gin.Context) {

	itemsCategoryLinkId, itemsCategoryLinkExists := c.Get(adg.ITEM_CATEGORY_LINK_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsCategoryLinkExists && userLoginExist {
		itemsCategoryLinkId := itemsCategoryLinkId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		itemCategoryLink, err := aasicl.CategoryLinkService.GetItemsCategoryLinkBySubCategoryPrivate(currentUser, itemsCategoryLinkId)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itemCategoryLink)
		return

	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemsCategoryLinkExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func InsertItemCategoryLinkPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")

	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		var itemCategoryLinkParams adm.CreateItemsCategoryLinkParams
		if err := c.ShouldBindJSON(&itemCategoryLinkParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := aasicl.CategoryLinkService.InsertItemsCategoryLinkPrivate(currentUser, itemCategoryLinkParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func UpdateItemCategoryLinkPrivate(c *gin.Context) {

	itemsCategoryLinkId, itemsCategoryLinkIdexists := c.Get(adg.ITEM_CATEGORY_LINK_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsCategoryLinkIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		iclId := itemsCategoryLinkId.(int32)

		updateMysqlItemCategoryLink := adm.UpdateItemsCategoryLinkParams{IclID: iclId}
		if err := c.ShouldBindJSON(&updateMysqlItemCategoryLink); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasicl.CategoryLinkService.UpdateItemsCategoryLinkPrivate(currentUser, updateMysqlItemCategoryLink)
		isErr := aac.EditErrorResponse(rows, iclId, err, c)

		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item category link %d updated ", iclId))
		}
		return

	}
	if !itemsCategoryLinkIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func DeleteItemCategoryLinkPrivate(c *gin.Context) {

	userLogin, userExists := c.Get("userLogin")
	itemsCategoryLinkId, itemCategoryLinkIdExists := c.Get(adg.ITEM_CATEGORY_ID)

	if userExists && itemCategoryLinkIdExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		itemCategoryLinkId := itemsCategoryLinkId.(int32)

		rows, err := aasicl.CategoryLinkService.DeleteItemsCategoryLinkPrivate(currentUser, itemCategoryLinkId)
		isErr := aac.EditErrorResponse(rows, itemCategoryLinkId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item category link %d deleted ", itemCategoryLinkId))
		}
		return

	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemCategoryLinkIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}
