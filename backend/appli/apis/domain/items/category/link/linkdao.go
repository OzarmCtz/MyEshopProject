package link

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListItemsCategoryLink() ([]adm.ItemsCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemCategoryLink, err := adm.QueriesDb.ListItemsCategoryLink(ctx)
	return itemCategoryLink, err
}

func ListItemsCategoryLinkByCategory(categoryId int32) ([]adm.ItemsCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemCategoryLink, err := adm.QueriesDb.ListItemsCategoryLinkByCategory(ctx, categoryId)
	return itemCategoryLink, err
}

func GetItemCategoryLinkBySubCategory(subCategoryId int32) (adm.ItemsCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemsCategoryLink, err := adm.QueriesDb.GetItemCategoryLinkBySubCategory(ctx, subCategoryId)
	return itemsCategoryLink, err
}

func GetItemsCategoryLink(isclId int32) (adm.ItemsCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemsCategoryLink, err := adm.QueriesDb.GetItemCategoryLink(ctx, isclId)
	return itemsCategoryLink, err
}

func InsertItemsCategoryLink(itemsCategoryLinkParams adm.CreateItemsCategoryLinkParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemsCategoryLink(ctx, itemsCategoryLinkParams)
	return res, err
}

func UpdateItemsCategoryLink(itemsCategoryLinkParams adm.UpdateItemsCategoryLinkParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemsCategoryLink(ctx, itemsCategoryLinkParams)
	return rows, err
}

func UpdateItemsCategoryLinkBySubCategory(itemsCategoryLinkParams adm.UpdateItemsCategoryLinkBySubCategoryParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemsCategoryLinkBySubCategory(ctx, itemsCategoryLinkParams)
	return rows, err
}

func DeleteItemsCategoryLink(isclId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItemsCategoryLink(ctx, isclId)
	return rows, err
}
