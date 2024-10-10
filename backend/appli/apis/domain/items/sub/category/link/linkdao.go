package link

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListItemsSubCategoryLink() ([]adm.ItemsSubCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemSubCategoryLink, err := adm.QueriesDb.ListItemsSubCategoryLink(ctx)
	return itemSubCategoryLink, err
}

func ListItemsBySubCategory(subCategoryId int32) ([]adm.ItemsSubCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.ListItemsSubCategoryLinkByCategory(ctx, subCategoryId)
	return res, err

}

func GetItemsSubCategoryLink(isclId int32) (adm.ItemsSubCategoryLink, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemsSubCategoryLink, err := adm.QueriesDb.GetItemsSubCategoryLink(ctx, isclId)
	return itemsSubCategoryLink, err
}

func GetItemsSubCategoryLinkByItem(isclId int32) (adm.ItemsSubCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	itemsSubCategoryLink, err := adm.QueriesDb.GetItemsSubCategoryLinkByItem(ctx, isclId)
	return itemsSubCategoryLink, err
}

func InsertItemsSubCategoryLinkBySubCategoryName(params adm.CreateItemsSubCategoryLinkBySubCategoryNameParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemsSubCategoryLinkBySubCategoryName(ctx, params)
	return res, err
}

func InsertItemsSubCategoryLink(itemsSubCategoryLinkParams adm.CreateItemsSubCategoryLinkParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItemsSubCategoryLink(ctx, itemsSubCategoryLinkParams)
	return res, err
}

func UpdateItemsSubCategoryLink(itemsSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemsSubCategoryLink(ctx, itemsSubCategoryLinkParams)
	return rows, err
}

func UpdateItemsSubCategoryByItemLink(itemsSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkByItemParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItemsSubCategoryLinkByItem(ctx, itemsSubCategoryLinkParams)
	return rows, err
}

func DeleteItemsSubCategoryLink(isclId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItemsSubCategoryLink(ctx, isclId)
	return rows, err
}
