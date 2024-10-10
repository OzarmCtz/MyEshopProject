package category

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListItemsSubCategory() ([]adm.ItemsSubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	ItemsSubCategory, err := adm.QueriesDb.ListItemsSubCategory(ctx)
	return ItemsSubCategory, err
}

func GetItemsSubCategory(subCategoryId int32) (adm.ItemsSubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemSubCategory, err := adm.QueriesDb.GetItemSubCategory(ctx, subCategoryId)
	return itemSubCategory, err
}

func GetItemsSubCategoryAndCategoryLinked(subCategoryId int32) (adm.GetItemSubCategoryAndCategoryLinkedRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemSubCategory, err := adm.QueriesDb.GetItemSubCategoryAndCategoryLinked(ctx, subCategoryId)
	return itemSubCategory, err
}

func ListItemsSubCategoryAndCategoryLinked() ([]adm.ListItemSubCategoryAndCategoryLinkedRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemSubCategory, err := adm.QueriesDb.ListItemSubCategoryAndCategoryLinked(ctx)
	return itemSubCategory, err
}

func InsertItemSubCategory(subCategoryParams adm.CreateItemSubCategoryParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemSubCategory(ctx, subCategoryParams)
	return res, err
}

func UpdateItemSubCategory(subCategoryParams adm.UpdateItemSubCategoryParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemSubCategory(ctx, subCategoryParams)
	return rows, err
}

func DeleteItemSubCategory(subCategoryId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItemSubCategory(ctx, subCategoryId)
	return rows, err
}
