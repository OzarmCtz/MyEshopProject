package link

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	aadd "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/discount"
	aadi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestDiscountLinkDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	insertDiscountParams := adm.CreateDiscountParams{
		DCode:        "TestCode",
		DDescription: adm.NullString{NullString: sql.NullString{String: "Test Description", Valid: true}},
		DStartTime:   adm.NullTime{NullTime: sql.NullTime{Time: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC), Valid: true}},
		DEndTime:     adm.NullTime{NullTime: sql.NullTime{Time: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC), Valid: true}},
		DZoneTime:    adm.NullString{NullString: sql.NullString{String: "UTC/PARIS", Valid: true}},
		DIsDisabled:  false,
		DPriceType:   "EUR",
		DValue:       10,
	}

	resInsertDiscount, err := aadd.InsertDiscount(insertDiscountParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	discountId, err := resInsertDiscount.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, discountId, int64(1))

	item, err := aadi.CreateItem()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	insertDiscountLinkParams := adm.CreateDiscountLinkParams{
		DlDiscountID: int32(discountId),
		DlItemsID:    adm.NullInt32{NullInt32: sql.NullInt32{Int32: int32(item.IID), Valid: true}},
	}

	resInsertDiscountLink, err := InsertDiscountLink(insertDiscountLinkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	discountLinkId, err := resInsertDiscountLink.LastInsertId()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, discountLinkId, int64(1))

	secondItem, err := aadi.CreateItem()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	updateDiscountLink, err := UpdateDiscountLink(adm.UpdateDiscountLinkParams{
		DlDiscountID: int32(discountId),
		DlItemsID:    adm.NullInt32{NullInt32: sql.NullInt32{Int32: int32(secondItem.IID), Valid: true}},
		DlID:         int32(discountLinkId),
	})

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, updateDiscountLink, int64(1))

	listDiscountsLinks, err := ListDiscountsLinks()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, len(listDiscountsLinks), 1)

	rows, err := DeleteDiscountLink(int32(discountLinkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

}
