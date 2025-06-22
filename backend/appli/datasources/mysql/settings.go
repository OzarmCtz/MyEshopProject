package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Db        *sql.DB
	QueriesDb *Queries
	DbName    string
)

// findEnvFile recherche le fichier .env en remontant dans l'arborescence
func findEnvFile() string {
	dir, _ := os.Getwd()

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath
		}

		// Remonter d'un niveau
		parent := filepath.Dir(dir)
		if parent == dir {
			break // Racine du système atteinte
		}
		dir = parent
	}

	log.Fatal("Could not find .env file")
	return ""
}

func MysqlInit(dbName string) {
	// Recherche automatique du fichier .env en remontant dans l'arborescence
	envPath := findEnvFile()

	isErr := godotenv.Load(envPath)
	if isErr != nil {
		log.Fatal("Error loading .env file:", isErr)
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
