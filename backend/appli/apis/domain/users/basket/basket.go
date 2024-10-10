package basket

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func GetUserBasket(basketId int32) (adm.UsersBasket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	basket, err := adm.QueriesDb.GetUserBasket(ctx, basketId)
	return basket, err
}

func ListUserBasketByUser(userId int32) ([]adm.UsersBasket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	basket, err := adm.QueriesDb.ListUserBasketByUser(ctx, userId)
	return basket, err
}

func InsertBasketList(basketParams adm.CreateUserBasketParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateUserBasket(ctx, basketParams)
	return res, err
}

func DeleteUserBasketByUserAndItem(deleteBasketParams adm.DeleteUserBasketByUserAndItemsParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteUserBasketByUserAndItems(ctx, deleteBasketParams)
	return rows, err
}
