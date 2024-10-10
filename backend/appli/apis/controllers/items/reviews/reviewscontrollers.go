package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aac "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/controllers"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	asir "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/reviews"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
)

func ListItemReviewByItemIdPublic(c *gin.Context) {

	reviewIdExists := c.Param(adg.ITEM_REVIEW_ID)

	if reviewIdExists != "" {
		reviewId, err := strconv.Atoi(reviewIdExists)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID format"})
			return
		}

		reviews, err := asir.ReviewsService.ListItemReviewByItemIdPublic(int32(reviewId))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(reviews) == 0 {
			c.JSON(http.StatusNotFound, reviews)
			return
		}

		c.JSON(http.StatusOK, reviews)
	}

	if reviewIdExists == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}
}

func ListItemReviewByUserIdPrivate(c *gin.Context) {
	userLogin, userExist := c.Get("userLogin")
	userReviewsId, userReviewIdExists := c.Get(adg.ITEM_REVIEW_ID)

	if userExist && userReviewIdExists {
		userReviewId := userReviewsId.(int32)
		currentUser := userLogin.(aadu.AppliUserLogin)
		reviews, err := asir.ReviewsService.ListItemReviewByUserIdPrivate(currentUser, userReviewId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(reviews) == 0 {
			c.JSON(http.StatusNotFound, reviews)
			return
		}

		c.JSON(http.StatusOK, reviews)

	}

	if !userExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !userReviewIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}

func ListItemReviewPrivate(c *gin.Context) {
	_, userExist := c.Get("userLogin")

	if userExist {

		items, err := asir.ReviewsService.ListItemReviewPrivate()

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

func InsertItemReviewPrivate(c *gin.Context) {

	userLogin, userLoginExist := c.Get("userLogin")
	if userLoginExist {
		currentUser := userLogin.(aadu.AppliUserLogin)

		reviewParams := adm.CreateItemReviewParams{}
		if err := c.ShouldBindJSON(&reviewParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		res, err := asir.ReviewsService.InsertItemReviewPrivate(currentUser, reviewParams)
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
func DeleteItemeReviewPrivate(c *gin.Context) {
	userLogin, userExists := c.Get("userLogin")
	reviewsId, reviewIdExists := c.Get(adg.ITEM_REVIEW_ID)

	if userExists && reviewIdExists {
		currentUser := userLogin.(aadu.AppliUserLogin)
		reviewId := reviewsId.(int32)

		rows, err := asir.ReviewsService.DeleteItemReviewByUserAndItemPrivate(currentUser, reviewId)
		isErr := aac.EditErrorResponse(rows, reviewId, err, c)
		if !isErr {
			c.JSON(http.StatusOK, fmt.Sprintf("item review %d deleted ", reviewId))
		}
		return

	}
	if !userExists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
	}

	if !reviewIdExists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf(adg.PARAM_ERROR_MESSAGE, c.Request.URL.Path)})
	}

}
