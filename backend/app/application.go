package app

import (
	"fmt"
	"log"
	"os"

	zaplogger "github.com/OzarmCtz/e_shop_backend_v1/app/utils/logger/zap"
	f "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/firebase"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router           = gin.Default()
	UsersByUID       = make(map[string]adm.User)
	MYSQL_APPLI_PORT string
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("MYSQL_APPLI_PORT : ", os.Getenv("MYSQL_APPLI_PORT"))
	MYSQL_APPLI_PORT = os.Getenv("MYSQL_APPLI_PORT")

	f.FirebaseInit()
	adm.MysqlInit(os.Getenv("MYSQL_APPLI_DB"))

}

func StartApp() {
	port := MYSQL_APPLI_PORT

	MapUrls()
	zaplogger.Info("Application started on port : " + port)
	router.Run(":" + port)

}
