package controllers

import (
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	"github.com/gin-gonic/gin"
)

func GetNoRowsErrorResponse(err error, c *gin.Context) bool {
	if err != nil {
		if err.Error() == adg.NO_ROWS_IN_RESULT_SET_ERROR {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return true
		}
		c.JSON(400, gin.H{"error": err.Error()})
	}
	return false
}

func EditErrorResponse(rows int64, id int32, err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return true
	}

	return false
}
