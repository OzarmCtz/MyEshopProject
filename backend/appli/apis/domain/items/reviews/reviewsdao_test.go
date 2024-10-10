package reviews

import (
	"fmt"
	"testing"

	aadi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items"
	aadu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	mysqldata "github.com/OzarmCtz/e_shop_backend_v1/appli/data/mysql"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aum "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/mysql"
	"github.com/go-playground/assert/v2"
)

func TestUserReviewDao(t *testing.T) {
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

	reviewInsert := adm.CreateItemReviewParams{
		IrUserID:   user.UID,
		IrItemsID:  item.IID,
		IrComments: "This is a test comment",
		IrStars:    5,
	}

	res, err := InsertReview(reviewInsert)
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

	review, err := GetUserReview(int32(resId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, review.IrUserID, user.UID)

	reviewByUserId, err := ListItemsReviewsByUser(user.UID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, len(reviewByUserId), 1)

	reviewByItemId, err := ListItemsReviewByItemId(item.IID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, len(reviewByItemId), 1)

	rows, err := DeleteItemReviewByUserAndItem(int32(resId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, rows, int64(1))
}
