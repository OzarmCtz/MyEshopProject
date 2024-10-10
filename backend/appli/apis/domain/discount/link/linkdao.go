package link

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListDiscountsLinks() ([]adm.DiscountLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	discountsLinks, err := adm.QueriesDb.ListDiscountsLinks(ctx)
	return discountsLinks, err
}

func GetDiscountLink(discountLinkId int32) (adm.DiscountLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	discountLink, err := adm.QueriesDb.GetDiscountLink(ctx, discountLinkId)
	return discountLink, err
}

func InsertDiscountLink(discountLinkParams adm.CreateDiscountLinkParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.CreateDiscountLink(ctx, discountLinkParams)
	return rows, err
}

func DeleteDiscountLink(discountLinkId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteDiscountLink(ctx, discountLinkId)
	return rows, err
}

func UpdateDiscountLink(discountLinkParams adm.UpdateDiscountLinkParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateDiscountLink(ctx, discountLinkParams)
	return rows, err
}
