package category

import (
	"database/sql"
	"fmt"
	"testing"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestItemCategoryDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	var category adm.ItemsCategory

	var CreateItemCategoryParams = adm.CreateItemCategoryParams{
		IcName:        "Category 1",
		IcDescription: adm.NullString{NullString: sql.NullString{String: "Description for Category 1", Valid: true}},
		IcPictureUrl:  adm.NullString{NullString: sql.NullString{String: "http://localhost:8080/picture.jpg", Valid: true}},
	}

	resInsertItemCategory, err := InsertItemCategory(CreateItemCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemCategoryId, err := resInsertItemCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	category, err = GetItemsCategory(int32(itemCategoryId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, category.IcName, "Category 1")

	var UpdateItemCategoryParams = adm.UpdateItemCategoryParams{
		IcName:        "Category 1",
		IcDescription: adm.NullString{NullString: sql.NullString{String: "Description for Category 1 Updated", Valid: true}},
		IcPictureUrl:  adm.NullString{NullString: sql.NullString{String: "http://localhost:8080/picture.jpg", Valid: true}},

		IcID: int32(category.IcID),
	}

	rows, err := UpdateItemCategory(UpdateItemCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	categoryList, err := ListItemsCategory()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, categoryList[0].IcDescription, adm.NullString{NullString: sql.NullString{String: "Description for Category 1 Updated", Valid: true}})

	rows, err = DeleteItemCategory(int32(category.IcID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))
}
