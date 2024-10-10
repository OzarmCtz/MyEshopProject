package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Db        *sql.DB
	QueriesDb *Queries
	DbName    string
)

func MysqlInit(dbName string) {

	isErr := godotenv.Load(".env")
	if isErr != nil {
		log.Fatal("Error loading .env file")
	}

	var dataSourceName string

	DbName = dbName

	if dbName != "" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			os.Getenv("MYSQL_USERNAME"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			dbName,
		)
	} else {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=true",
			os.Getenv("MYSQL_USERNAME"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
		)
	}

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Connection to the mysql DB couldn't be established, error:%v", err)
	}
	QueriesDb = New(Db)
	if err = Db.Ping(); err != nil {
		log.Fatalf("Connection to the mysql DB couldn't be established, error: %v", err)
	}

	log.Println("Mysql Database successfully configured")
}
