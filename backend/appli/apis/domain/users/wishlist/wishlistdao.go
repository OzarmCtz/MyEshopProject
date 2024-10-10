package wishlist

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func GetUserWishList(wishListID int32) (adm.UsersWishlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	wishList, err := adm.QueriesDb.GetUserWishList(ctx, wishListID)
	return wishList, err
}

func ListUserWishListByUser(userId int32) ([]adm.UsersWishlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	wishList, err := adm.QueriesDb.ListUserWishListByUser(ctx, userId)
	return wishList, err
}

func InsertUserWishList(wishListParams adm.CreateUserWishListParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateUserWishList(ctx, wishListParams)
	return res, err
}

func DeleteUserWishListByUserAndItem(deleteUserWishListParams adm.DeleteUserWishListByUserAndItemsParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteUserWishListByUserAndItems(ctx, deleteUserWishListParams)
	return rows, err
}
