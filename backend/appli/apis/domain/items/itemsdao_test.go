package items

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestItemDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	// Create Item

	item, err := CreateItem()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, item.IID, int32(item.IID))

	itemUpdateItems := adm.UpdateItemParams{
		ITitle:       "Item 1",
		IDescription: adm.NullString{NullString: sql.NullString{String: "Description for Item 1 Updated", Valid: true}},
		IPrice:       "0.0015",
		IQuantity:    adm.NullInt32{NullInt32: sql.NullInt32{Int32: 1, Valid: true}},
		IPictureUrl:  adm.NullString{NullString: sql.NullString{String: " https://domain.com/images/items1.png", Valid: true}},
		IFilePath:    adm.NullString{NullString: sql.NullString{String: "https://domain.com/files/items1.pdf", Valid: true}},
		IIsDisabled:  false,
		IReleaseDate: time.Now(),
		IID:          int32(item.IID),
	}

	rows, err := UpdateItem(itemUpdateItems)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	items, err := ListItems()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, items[0].IDescription, adm.NullString{NullString: sql.NullString{String: "Description for Item 1 Updated", Valid: true}})

	itemByActivity, err := ListItemByActivity(false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, itemByActivity[0].IIsDisabled, false)

	rows, err = DeleteItem(int32(item.IID))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

}
