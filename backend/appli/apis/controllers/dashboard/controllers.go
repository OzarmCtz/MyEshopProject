package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccess(c *gin.Context) {
	c.JSON(http.StatusOK, 1)
}
