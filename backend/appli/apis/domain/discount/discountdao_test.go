package discount

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

func TestDiscountDao(t *testing.T) {
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

	resInsertDiscount, err := InsertDiscount(insertDiscountParams)
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

	getDiscount, err := GetDiscount(int32(discountId))

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, getDiscount.DCode, "TestCode")

	getDiscountByCode, err := GetDiscountByCode("TestCode")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, getDiscountByCode.DDescription, adm.NullString{NullString: sql.NullString{String: "Test Description", Valid: true}})

	updateDiscountParams := adm.UpdateDiscountParams{
		DCode:        "TestCode",
		DDescription: adm.NullString{NullString: sql.NullString{String: "Test Description Updated", Valid: true}},
		DStartTime:   adm.NullTime{NullTime: sql.NullTime{Time: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC), Valid: true}},
		DEndTime:     adm.NullTime{NullTime: sql.NullTime{Time: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC), Valid: true}},
		DZoneTime:    adm.NullString{NullString: sql.NullString{String: "UTC/PARIS", Valid: true}},
		DIsDisabled:  false,
		DPriceType:   "EUR",
		DValue:       10,
		DID:          int32(discountId),
	}

	rows, err := UpdateDiscount(updateDiscountParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

	discounts, err := ListDiscounts()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, discounts[0].DDescription, adm.NullString{NullString: sql.NullString{String: "Test Description Updated", Valid: true}})

	rows, err = DeleteDiscount(int32(discountId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))
}
