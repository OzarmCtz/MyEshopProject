package categorycontroller

import (
	"fmt"
	"net/http"
	"strconv"

	aasicl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/category/link"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/sub/category"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemSubCategoryPublic(c *gin.Context) {
	items, err := aasisc.ItemsSubCategoryService.ListItemsSubCategoryPublic()

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

func ListItemSubCategoryPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		items, err := aasisc.ItemsSubCategoryService.ListItemsSubCategoryAndCategoryLinkedPrivate(currentUser)

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

func GetItemSubCategoryPublic(c *gin.Context) {

	itemSubCategoryIdExists := c.Param(adg.ITEM_SUB_CATEGORY_ID)

	if itemSubCategoryIdExists != "" {
		itemSubCategoryId, err := strconv.Atoi(itemSubCategoryIdExists)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item sub category ID format"})
			return
		}
		itemCategory, err := aasisc.ItemsSubCategoryService.GetItemsSubCategoryPublic(int32(itemSubCategoryId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itemCategory)
		return
	}

	if itemSubCategoryIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func GetItemSubCategoryPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")
	itemSubCategory, itemSubCategoryIdExists := c.Get(adg.ITEM_SUB_CATEGORY_ID)

	if itemSubCategoryIdExists && userLoginExist {
		itemSubCategoryId := itemSubCategory.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)

		itemCategory, err := aasisc.ItemsSubCategoryService.GetItemsSubCategoryAndCategoryLinkedPrivate(currentUser, itemSubCategoryId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, itemCategory)
		return
	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemSubCategoryIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func InsertSubItemCategoryPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")

	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		var itemSubCategoryWithCategoryLinkParams aum.CreateItemSubCategoryWithCategoryLinkParams
		if err := c.ShouldBindJSON(&itemSubCategoryWithCategoryLinkParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		itemSubCategoryParams := adm.CreateItemSubCategoryParams{
			IscName:        itemSubCategoryWithCategoryLinkParams.IscName,
			IscDescription: itemSubCategoryWithCategoryLinkParams.IscDescription,
			IscPictureUrl:  itemSubCategoryWithCategoryLinkParams.IscPictureUrl,
		}

		res, err := aasisc.ItemsSubCategoryService.InserItemsSubCategoryPrivate(currentUser, itemSubCategoryParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		itemCategoryLinkParams := adm.CreateItemsCategoryLinkParams{
			IclItemsSubCategoryID: res.IscID,
			IclItemsCategoryID:    itemSubCategoryWithCategoryLinkParams.IclItemsCategoryID,
		}

		result, err := aasicl.CategoryLinkService.InsertItemsCategoryLinkPrivate(currentUser, itemCategoryLinkParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		responseArray := gin.H{
			"itemSubCategory":  res,
			"itemCategoryLink": result,
		}

		c.JSON(http.StatusOK, &responseArray)
		return
	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func UpdateItemSubCategoryPrivate(c *gin.Context) {

	itemsSubCategoryId, itemsSubCategoryIdexists := c.Get(adg.ITEM_SUB_CATEGORY_ID)

	userLogin, userLoginExist := c.Get("userLogin")

	if itemsSubCategoryIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		iscId := itemsSubCategoryId.(int32)

		updateMysqlItemSubCategory := aum.UpdateItemSubCategoryWithCategoryLinkParams{IscID: iscId}
		if err := c.ShouldBindJSON(&updateMysqlItemSubCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateItemSubCategoryParams := adm.UpdateItemSubCategoryParams{
			IscName:        updateMysqlItemSubCategory.IscName,
			IscDescription: updateMysqlItemSubCategory.IscDescription,
			IscPictureUrl:  updateMysqlItemSubCategory.IscPictureUrl,
			IscID:          updateMysqlItemSubCategory.IscID,
		}

		rows, err := aasisc.ItemsSubCategoryService.UpdateItemsSubCategoryPrivate(currentUser, updateItemSubCategoryParams)
		isErr := aac.EditErrorResponse(rows, iscId, err, c)

		updateItemSubCategoryLinkParams := adm.UpdateItemsCategoryLinkBySubCategoryParams{
			IclItemsSubCategoryID: iscId,
			IclItemsCategoryID:    updateMysqlItemSubCategory.IclItemsCategoryID,
		}

		result, err := aasicl.CategoryLinkService.UpdateItemsCategoryLinkBySubCategoryPrivate(currentUser, updateItemSubCategoryLinkParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		isErrLink := aac.EditErrorResponse(result, iscId, err, c)

		if !isErr && !isErrLink {
			c.JSON(http.StatusOK, fmt.Sprintf("item %d updated ", iscId))
		}
		return

	}
	if !itemsSubCategoryIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func DeleteItemSubCategoryPrivate(c *gin.Context) {

	userLogin, userExists := c.Get("userLogin")
	itemsSubCategoryId, itemSubCategoryIdExists := c.Get(adg.ITEM_SUB_CATEGORY_ID)

	if userExists && itemSubCategoryIdExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		itemSubCategoryId := itemsSubCategoryId.(int32)

		rows, err := aasisc.ItemsSubCategoryService.DeleteItemsSubCategoryPrivate(currentUser, itemSubCategoryId)
		isErr := aac.EditErrorResponse(rows, itemSubCategoryId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item %d deleted ", itemSubCategoryId))
		}
		return

	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemSubCategoryIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}
