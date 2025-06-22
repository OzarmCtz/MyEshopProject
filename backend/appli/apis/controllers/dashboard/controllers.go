package dashboard

import (
	"net/http"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasd "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/dashboard"
	"github.com/gin-gonic/gin"
)

func GetAccess(c *gin.Context) {
	userLogin, userLoginExist := c.Get("userLogin")
	if !userLoginExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": adg.USER_ERROR_MESSAGE})
		return
	}

	currentUser := userLogin.(aadu.AppliUserLogin)
	err := aasd.DashboardService.CheckLoginDashboardAccess(currentUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// L'utilisateur est autorisé, renvoyer status 200
	c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
}
