package linkcontroller

import (
	"fmt"
	"net/http"
	"strconv"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasiscl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/sub/category/link"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemsBySubCategoryPublic(c *gin.Context) {

	itemSubCategoryLinkIdExists := c.Param(adg.ITEM_SUB_CATEGORY_LINK_ID)

	if itemSubCategoryLinkIdExists != "" {
		itemSubCategoryId, err := strconv.Atoi(itemSubCategoryLinkIdExists)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item category ID format"})
			return
		}

		items, err := aasiscl.SubCategoryLinkService.ListItemsBySubCategoryPublic(int32(itemSubCategoryId))

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

	if itemSubCategoryLinkIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func ListItemSubCategoryLinkPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		itemsSubCategoryLink, err := aasiscl.SubCategoryLinkService.ListItemsSubCategoryLinkPrivate(currentUser)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(itemsSubCategoryLink) == 0 {
			c.JSON(http.StatusNotFound, itemsSubCategoryLink)
			return
		}

		c.JSON(http.StatusOK, itemsSubCategoryLink)
		return

	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

}

func GetItemSubCategoryLinkPrivate(c *gin.Context) {

	itemsSubCategoryLinkId, itemsSubCategoryLinkExists := c.Get(adg.ITEM_SUB_CATEGORY_LINK_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsSubCategoryLinkExists && userLoginExist {
		itemsSubCategoryLinkId := itemsSubCategoryLinkId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		// return item sub category link by item
		itemSubCategoryLink, err := aasiscl.SubCategoryLinkService.GetItemsSubCategoryLinkByItemPrivate(currentUser, itemsSubCategoryLinkId)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itemSubCategoryLink)
		return

	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemsSubCategoryLinkExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func InsertItemSubCategoryLinkPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")

	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		var itemSubCategoryLinkParams adm.CreateItemsSubCategoryLinkParams
		if err := c.ShouldBindJSON(&itemSubCategoryLinkParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := aasiscl.SubCategoryLinkService.InsertItemsSubCategoryLinkPrivate(currentUser, itemSubCategoryLinkParams)
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

	itemsSubCategoryLinkId, itemsSubCategoryLinkIdexists := c.Get(adg.ITEM_SUB_CATEGORY_LINK_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsSubCategoryLinkIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		isclId := itemsSubCategoryLinkId.(int32)

		updateMysqlItemCategoryLink := adm.UpdateItemsSubCategoryLinkParams{IsclID: isclId}
		if err := c.ShouldBindJSON(&updateMysqlItemCategoryLink); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rows, err := aasiscl.SubCategoryLinkService.UpdateSubItemsCategoryLinkPrivate(currentUser, updateMysqlItemCategoryLink)
		isErr := aac.EditErrorResponse(rows, isclId, err, c)

		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item sub category link %d updated ", isclId))
		}
		return

	}
	if !itemsSubCategoryLinkIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func DeleteItemCategoryLinkPrivate(c *gin.Context) {

	userLogin, userExists := c.Get("userLogin")
	itemsSubCategoryLinkId, itemSubCategoryLinkIdExists := c.Get(adg.ITEM_SUB_CATEGORY_LINK_ID)

	if userExists && itemSubCategoryLinkIdExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		itemSubCategoryLinkId := itemsSubCategoryLinkId.(int32)

		rows, err := aasiscl.SubCategoryLinkService.DeleteItemsSunCategoryLinkPrivate(currentUser, itemSubCategoryLinkId)
		isErr := aac.EditErrorResponse(rows, itemSubCategoryLinkId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item sub category link %d deleted ", itemSubCategoryLinkId))
		}
		return

	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemSubCategoryLinkIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}
