package wishlist

import (
	"database/sql"
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

	wishListInsert := adm.CreateUserWishListParams{
		WlUserID:     user.UID,
		WlItemsID:    adm.NullInt32{NullInt32: sql.NullInt32{Int32: item.IID, Valid: true}},
		WlTimesAdded: time.Now(),
	}

	res, err := InsertUserWishList(wishListInsert)
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

	wishList, err := GetUserWishList(int32(resId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, wishList.WlUserID, user.UID)

	wishListByUser, err := ListUserWishListByUser(user.UID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, wishListByUser[0].WlUserID, user.UID)

	deleteWishListParams := adm.DeleteUserWishListByUserAndItemsParams{
		WlUserID:  user.UID,
		WlItemsID: adm.NullInt32{NullInt32: sql.NullInt32{Int32: item.IID, Valid: true}},
	}

	rows, err := DeleteUserWishListByUserAndItem(deleteWishListParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))

}
