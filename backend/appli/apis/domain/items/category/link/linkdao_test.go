package link

import (
	"database/sql"
	"fmt"
	"testing"

	adic "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/category"
	adisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/sub/category"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestItemDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	var CreateItemCategoryParams = adm.CreateItemCategoryParams{
		IcName:        "Category 1",
		IcDescription: adm.NullString{NullString: sql.NullString{String: "Description for Category 1", Valid: true}},
	}

	resInsertItemCategory, err := adic.InsertItemCategory(CreateItemCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemCategoryId, err := resInsertItemCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var CreateItemSubCategoryParams = adm.CreateItemSubCategoryParams{
		IscName:        "Category 1",
		IscDescription: adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 1", Valid: true}},
	}

	resInsertItemSubCategory, err := adisc.InsertItemSubCategory(CreateItemSubCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSubCategoryId, err := resInsertItemSubCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var CreateItemsCategoryLinkParams = adm.CreateItemsCategoryLinkParams{
		IclItemsSubCategoryID: int32(itemSubCategoryId),
		IclItemsCategoryID:    int32(itemCategoryId),
	}

	createItemsCategoryLink, err := InsertItemsCategoryLink(CreateItemsCategoryLinkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemsCategoryLinkID, err := createItemsCategoryLink.LastInsertId()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemsCategoryLink, err := GetItemsCategoryLink(int32(itemsCategoryLinkID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, itemsCategoryLink.IclItemsCategoryID, int32(itemCategoryId))

	var CreateSecondItemCategoryParams = adm.CreateItemCategoryParams{
		IcName:        "Category 2",
		IcDescription: adm.NullString{NullString: sql.NullString{String: "Description for Category 2", Valid: true}},
	}

	resInsertSecondItemCategory, err := adic.InsertItemCategory(CreateSecondItemCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSecondCategoryId, err := resInsertSecondItemCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var UpdateItemsCategoryLinkParams = adm.UpdateItemsCategoryLinkParams{
		IclItemsCategoryID:    int32(itemSecondCategoryId),
		IclItemsSubCategoryID: int32(itemCategoryId),
		IclID:                 int32(itemsCategoryLinkID),
	}

	rows, err := UpdateItemsCategoryLink(UpdateItemsCategoryLinkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	listItemsCategoryLink, err := ListItemsCategoryLink()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, listItemsCategoryLink[0].IclItemsCategoryID, int32(itemSecondCategoryId))

	rows, err = DeleteItemsCategoryLink(int32(itemsCategoryLinkID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

}
