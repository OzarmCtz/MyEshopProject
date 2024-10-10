package mysql

import (
	"database/sql"
	"fmt"
	"log"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func DeleteDatasFromDB(dbTables []string, db *sql.DB) {
	for _, dbTable := range dbTables {
		query := fmt.Sprintf("DELETE FROM `%s`", dbTable)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Datas from all tables were deleted successfully\n")
}

func ResetDbAutoIncrement(dbTables []string, db *sql.DB) {
	for _, dbTable := range dbTables {
		query := fmt.Sprintf("ALTER TABLE `%s`AUTO_INCREMENT = 1", dbTable)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetQueriesDbFromDbPool(clientId int32) *adm.Queries {
	return adg.QueriesUserDbPool[clientId]
}
