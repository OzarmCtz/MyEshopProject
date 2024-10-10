package itemscontroller

import (
	"fmt"
	"net/http"
	"strconv"

	aasiscl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/sub/category/link"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemPublic(c *gin.Context) {
	items, err := aasi.ItemsService.ListItemsPublic()

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

func ListItemPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")

	if userExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		items, err := aasi.ItemsService.ListItemsPrivate(currentUser)

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
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}

func GetItemPublic(c *gin.Context) {

	itemIdExists := c.Param(adg.ITEM_ID)

	if itemIdExists != "" {
		itemId, err := strconv.Atoi(itemIdExists)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
			return
		}
		item, err := aasi.ItemsService.GetItemsPublic(int32(itemId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, item)

	}

	if itemIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func GetItemPrivate(c *gin.Context) {

	itemsId, itemIdExists := c.Get(adg.ITEM_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if itemIdExists && userLoginExist {
		itemId := itemsId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		item, err := aasi.ItemsService.GetItemsPrivate(currentUser, itemId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, item)

	}

	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !itemIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func DeleteItemPrivate(c *gin.Context) {
	itemsId, itemIdExists := c.Get(adg.ITEM_ID)
	userLogin, userExists := c.Get("userLogin")

	if itemIdExists && userExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		iId := itemsId.(int32)
		rows, err := aasi.ItemsService.DeleteItems(currentUser, iId)
		isErr := aac.EditErrorResponse(rows, iId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item %d deleted ", iId))
		}
		return
	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !itemIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func UpdateItemPrivate(c *gin.Context) {
	itemsId, itemsIdexists := c.Get(adg.ITEM_ID)
	userLogin, userLoginExist := c.Get("userLogin")
	if itemsIdexists && userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)
		iId := itemsId.(int32)
		updateMysqlItem := aum.UpdateItemWithSubCategoryParams{IID: iId}
		if err := c.ShouldBindJSON(&updateMysqlItem); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		updateItemParams := adm.UpdateItemParams{
			ITitle:       updateMysqlItem.ITitle,
			IDescription: updateMysqlItem.IDescription,
			IPrice:       updateMysqlItem.IPrice,
			IQuantity:    updateMysqlItem.IQuantity,
			IPictureUrl:  updateMysqlItem.IPictureUrl,
			IFilePath:    updateMysqlItem.IFilePath,
			IIsDisabled:  updateMysqlItem.IIsDisabled,
			IID:          updateMysqlItem.IID,
		}

		rows, err := aasi.ItemsService.UpdateItems(currentUser, updateItemParams)
		isErr := aac.EditErrorResponse(rows, iId, err, c)

		subCategoryLink := adm.UpdateItemsSubCategoryLinkByItemParams{
			IsclSubCategoryID: updateMysqlItem.IsclSubCategoryID,
			IsclItemsID:       updateMysqlItem.IID,
		}

		resSubCategoryLink, err := aasiscl.SubCategoryLinkService.UpdateSubItemsCategoryByItemLinkPrivate(currentUser, subCategoryLink)
		isErrSubCategory := aac.EditErrorResponse(resSubCategoryLink, updateMysqlItem.IsclSubCategoryID, err, c)

		if !isErr && !isErrSubCategory {
			c.JSON(http.StatusOK, fmt.Sprintf("item %d updated ", iId))
		}
		return
	}
	if !itemsIdexists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}
}

func InsertItemPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")
	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		itemSubCategoryParams := aum.CreateItemWithSubCategoryParams{}
		if err := c.ShouldBindJSON(&itemSubCategoryParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		itemParams := adm.CreateItemParams{
			ITitle:       itemSubCategoryParams.ITitle,
			IDescription: itemSubCategoryParams.IDescription,
			IPrice:       itemSubCategoryParams.IPrice,
			IQuantity:    itemSubCategoryParams.IQuantity,
			IPictureUrl:  itemSubCategoryParams.IPictureUrl,
			IFilePath:    itemSubCategoryParams.IFilePath,
			IIsDisabled:  itemSubCategoryParams.IIsDisabled,
			IReleaseDate: itemSubCategoryParams.IReleaseDate,
		}

		res, err := aasi.ItemsService.InsertItems(currentUser, itemParams)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		subCategoryLink := adm.CreateItemsSubCategoryLinkParams{
			IsclItemsID:       res.IID,
			IsclSubCategoryID: itemSubCategoryParams.IsclSubCategoryID,
		}

		resSubCategoryLink, err := aasiscl.SubCategoryLinkService.InsertItemsSubCategoryLinkPrivate(currentUser, subCategoryLink)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		responseArray := gin.H{
			"item":                res,
			"itemSubCategoryLink": resSubCategoryLink,
		}

		c.JSON(http.StatusCreated, &responseArray)
		return

	}

	if !userLoginExist {
		c.JSON(http.StatusNotFound, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

}
