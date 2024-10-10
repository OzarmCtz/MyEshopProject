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

	var subCategory adm.ItemsSubCategory

	var CreateItemSubCategoryParams = adm.CreateItemSubCategoryParams{
		IscName:        "Category 1",
		IscDescription: adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 1", Valid: true}},
	}

	resInsertItemSubCategory, err := InsertItemSubCategory(CreateItemSubCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSubCategoryId, err := resInsertItemSubCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	subCategory, err = GetItemsSubCategory(int32(itemSubCategoryId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, subCategory.IscName, "Category 1")

	var UpdateItemCategoryParams = adm.UpdateItemSubCategoryParams{
		IscName:        "Category 1",
		IscDescription: adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 1 Updated", Valid: true}},
		IscID:          int32(subCategory.IscID),
	}

	rows, err := UpdateItemSubCategory(UpdateItemCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	categoryList, err := ListItemsSubCategory()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, categoryList[0].IscDescription, adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 1 Updated", Valid: true}})

	rows, err = DeleteItemSubCategory(int32(subCategory.IscID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))
}
