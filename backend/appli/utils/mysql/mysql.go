package mysqlutils

import (
	"database/sql"
	"time"

	adg "github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	aum "github.com/OzarmCtz/e_shop_backend_v1/app/utils/mysql"
	adtm "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func setUserDbPools() {
	adg.MysqlUserDbPool[adg.NO_CLIENT_QUERIES_DB_ID] = adm.Db
	adg.QueriesUserDbPool[adg.NO_CLIENT_QUERIES_DB_ID] = adm.New(adm.Db)
}

func InitAppliTestDB(db *sql.DB) {
	CleanDatabase(db)
	setUserDbPools()
	// GenerateGroupsAndPrivileges()
}

func CleanDatabase(db *sql.DB) {
	DeleteDatasFromAppliDB(db)
	ResetAppliDbAutoIncrement(db)
}

func ResetAppliDbAutoIncrement(db *sql.DB) {
	aum.ResetDbAutoIncrement(adtm.AppliDbTables, db)
}

func DeleteDatasFromAppliDB(db *sql.DB) {
	aum.DeleteDatasFromDB(adtm.AppliDbTables, db)
}

func GetQueriesDbFromDbPool(clientId int32) *adm.Queries {
	return adg.QueriesUserDbPool[clientId]
}

type CreateItemSubCategoryWithCategoryLinkParams struct {
	IscName            string         `json:"isc_name"`
	IscDescription     adm.NullString `json:"isc_description"`
	IscPictureUrl      adm.NullString `json:"isc_picture_url"`
	IclItemsCategoryID int32          `json:"icl_items_category_id"`
}

type UpdateItemSubCategoryWithCategoryLinkParams struct {
	IscID              int32          `json:"isc_id"`
	IscName            string         `json:"isc_name"`
	IscDescription     adm.NullString `json:"isc_description"`
	IscPictureUrl      adm.NullString `json:"isc_picture_url"`
	IclItemsCategoryID int32          `json:"icl_items_category_id"`
}
type CreateItemWithSubCategoryParams struct {
	ITitle            string         `json:"i_title"`
	IDescription      adm.NullString `json:"i_description"`
	IPrice            string         `json:"i_price"`
	IQuantity         adm.NullInt32  `json:"i_quantity"`
	IPictureUrl       adm.NullString `json:"i_picture_url"`
	IFilePath         adm.NullString `json:"i_file_path"`
	IIsDisabled       bool           `json:"i_is_disabled"`
	IReleaseDate      time.Time      `json:"i_release_date"`
	IsclSubCategoryID int32          `json:"iscl_sub_category_id"`
}

type UpdateItemWithSubCategoryParams struct {
	IID               int32          `json:"i_id"`
	ITitle            string         `json:"i_title"`
	IDescription      adm.NullString `json:"i_description"`
	IPrice            string         `json:"i_price"`
	IQuantity         adm.NullInt32  `json:"i_quantity"`
	IPictureUrl       adm.NullString `json:"i_picture_url"`
	IFilePath         adm.NullString `json:"i_file_path"`
	IIsDisabled       bool           `json:"i_is_disabled"`
	IReleaseDate      time.Time      `json:"i_release_date"`
	IsclSubCategoryID int32          `json:"iscl_sub_category_id"`
}
