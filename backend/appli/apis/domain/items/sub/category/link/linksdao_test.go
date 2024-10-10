package link

import (
	"database/sql"
	"fmt"
	"testing"

	aadi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items"
	aadisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/sub/category"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestItemDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	item, err := aadi.CreateItem()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var CreateItemSubCategoryParams = adm.CreateItemSubCategoryParams{
		IscName:        "Category 1",
		IscDescription: adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 1", Valid: true}},
	}

	resInsertItemSubCategory, err := aadisc.InsertItemSubCategory(CreateItemSubCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSubCategoryId, err := resInsertItemSubCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var CreateItemSubCategoryLinkParams = adm.CreateItemsSubCategoryLinkParams{
		IsclItemsID:       int32(item.IID),
		IsclSubCategoryID: int32(itemSubCategoryId),
	}

	resInsertItemSubCategoryLink, err := InsertItemsSubCategoryLink(CreateItemSubCategoryLinkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSubCategoryLinkId, err := resInsertItemSubCategoryLink.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, itemSubCategoryLinkId, int64(itemSubCategoryLinkId))

	var CreateSecondItemSubCategoryParams = adm.CreateItemSubCategoryParams{
		IscName:        "Category 2",
		IscDescription: adm.NullString{NullString: sql.NullString{String: "Description for Sub Category 2", Valid: true}},
	}

	resInsertSecondItemSubCategory, err := aadisc.InsertItemSubCategory(CreateSecondItemSubCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itemSecondSubCategoryId, err := resInsertSecondItemSubCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	var UpdateItemSubCategoryLinkParams = adm.UpdateItemsSubCategoryLinkParams{
		IsclItemsID:       int32(item.IID),
		IsclSubCategoryID: int32(itemSecondSubCategoryId),
		IsclID:            int32(itemSubCategoryLinkId),
	}

	rows, err := UpdateItemsSubCategoryLink(UpdateItemSubCategoryLinkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	itemSubCategoryLink, err := GetItemsSubCategoryLink(int32(itemSubCategoryLinkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, itemSubCategoryLink.IsclItemsID, int32(item.IID))

	itemSubCategoryLinkList, err := ListItemsSubCategoryLink()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, itemSubCategoryLinkList[0].IsclItemsID, int32(item.IID))

	_, err = DeleteItemsSubCategoryLink(int32(itemSubCategoryLinkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
