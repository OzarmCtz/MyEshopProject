package basket

import (
	"fmt"
	"testing"
	"time"

	aadi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestUserWishListDao(t *testing.T) {
	adm.MysqlInit(mysqldata.MYSQL_APPLI_TEST_DB)
	aum.InitAppliTestDB(adm.Db)

	user, err := aadu.CreateUser()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	item, err := aadi.CreateItem()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	basketInsert := adm.CreateUserBasketParams{
		UbUserID:    user.UID,
		UbItemsID:   item.IID,
		UbTimeAdded: time.Now(),
		UbQuantity:  1,
	}

	res, err := InsertBasketList(basketInsert)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, resId, int64(1))

	basket, err := GetUserBasket(int32(resId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, basket.UbUserID, user.UID)

	basketByUser, err := ListUserBasketByUser(user.UID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, basketByUser[0].UbUserID, user.UID)

	deleteBasketParams := adm.DeleteUserBasketByUserAndItemsParams{
		UbUserID:  user.UID,
		UbItemsID: item.IID,
	}

	rows, err := DeleteUserBasketByUserAndItem(deleteBasketParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

}
