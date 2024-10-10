package discount

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListDiscounts() ([]adm.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	discounts, err := adm.QueriesDb.ListDiscounts(ctx)
	return discounts, err
}

func GetDiscount(discountId int32) (adm.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	discount, err := adm.QueriesDb.GetDiscount(ctx, discountId)
	return discount, err
}

func GetDiscountByCode(discountCode string) (adm.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	discount, err := adm.QueriesDb.GetDiscountByCode(ctx, discountCode)
	return discount, err
}

func InsertDiscount(discountParams adm.CreateDiscountParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.CreateDiscount(ctx, discountParams)
	return rows, err
}

func DeleteDiscount(discountId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteDiscount(ctx, discountId)
	return rows, err
}

func UpdateDiscount(discountParams adm.UpdateDiscountParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateDiscount(ctx, discountParams)
	return rows, err
}
