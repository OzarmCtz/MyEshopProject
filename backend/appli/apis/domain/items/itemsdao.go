package items

import (
	"context"
	"database/sql"

	"github.com/OzarmCtz/e_shop_backend_v1/app/data/global"
	adt "github.com/OzarmCtz/e_shop_backend_v1/appli/data/tests"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

func ListItems() ([]adm.ListItemsRow, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	items, err := adm.QueriesDb.ListItems(ctx)
	return items, err
}

func GetItems(itemID int32) (adm.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	item, err := adm.QueriesDb.GetItem(ctx, itemID)
	return item, err
}

func InsertItem(itemParams adm.CreateItemParams) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	res, err := adm.QueriesDb.CreateItem(ctx, itemParams)
	return res, err
}

func UpdateItem(itemParams adm.UpdateItemParams) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.UpdateItem(ctx, itemParams)
	return rows, err
}

func DeleteItem(itemId int32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	rows, err := adm.QueriesDb.DeleteItem(ctx, itemId)
	return rows, err
}

func ListItemByActivity(isDisabled bool) ([]adm.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.Timeout)
	defer cancel()
	items, err := adm.QueriesDb.ListActiveItems(ctx, isDisabled)
	return items, err
}

// Utility function
func CreateItem() (adm.Item, error) {
	var item adm.Item

	resInsertItem, err := InsertItem(adt.CreateItemParams)
	if err != nil {
		return item, nil
	}

	itemId, err := resInsertItem.LastInsertId()
	if err != nil {
		return item, err
	}

	item, err = GetItems(int32(itemId))
	if err != nil {
		return item, err
	}

	return item, nil
}
