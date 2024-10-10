package category

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListItemsCategory() ([]adm.ItemsCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	ItemsCategory, err := adm.QueriesDb.ListItemsCategory(ctx)
	return ItemsCategory, err
}

func ListItemsCategoryAndOccurence() ([]adm.ListItemsCategoryAndOccurenceRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	ItemsCategory, err := adm.QueriesDb.ListItemsCategoryAndOccurence(ctx)
	return ItemsCategory, err
}

func GetItemsCategory(categoryId int32) (adm.ItemsCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemCategory, err := adm.QueriesDb.GetItemCategory(ctx, categoryId)
	return itemCategory, err
}

func InsertItemCategory(categoryParams adm.CreateItemCategoryParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemCategory(ctx, categoryParams)
	return res, err
}

func UpdateItemCategory(categoryParams adm.UpdateItemCategoryParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemCategory(ctx, categoryParams)
	return rows, err
}

func DeleteItemCategory(categoryId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItemCategory(ctx, categoryId)
	return rows, err
}
