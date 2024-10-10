package link

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"

	aadiscl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/sub/category/link"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items"
	aasisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/sub/category"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	SubCategoryLinkService subCategoryLinkServiceInterface = &subCategoryLinkService{}
)

type ItemsSubCategoryLink []adm.ItemsSubCategoryLink
type ItemSubCategoryLink adm.ItemsSubCategoryLink
type subCategoryLinkService struct{}

type subCategoryLinkServiceInterface interface {
	ListItemsBySubCategoryPublic(subCategoryId int32) ([]aasi.ItemPbResponse, error)
	UpdateSubItemsCategoryByItemLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkByItemParams) (int64, error)
	GetItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (adm.ItemsSubCategoryLink, error)
	InsertItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.CreateItemsSubCategoryLinkParams) (adm.ItemsSubCategoryLink, error)
	UpdateSubItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkParams) (int64, error)
	DeleteItemsSunCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (int64, error)
	ListItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin) ([]adm.ItemsSubCategoryLink, error)
	GetItemsSubCategoryLinkByItemPrivate(currentUser adu.AppliUserLogin, itemId int32) (aasisc.ItemSubCategoryRp, error)
}

// ListItemsBySubCategoryPublic is a method to list items by sub category
func (scls *subCategoryLinkService) ListItemsBySubCategoryPublic(subCategoryId int32) ([]aasi.ItemPbResponse, error) {

	rawItems, err := aadiscl.ListItemsBySubCategory(subCategoryId)
	if err != nil {
		return nil, err
	}

	publicItems := make([]aasi.ItemPbResponse, 0, len(rawItems))

	for _, item := range rawItems {
		publicItem, err := aasi.ItemsService.GetItemsPublic(item.IsclItemsID)
		if err != nil {
			return nil, err
		}
		publicItems = append(publicItems, publicItem)

	}

	return publicItems, nil

}

func (scls *subCategoryLinkService) GetItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (adm.ItemsSubCategoryLink, error) {

	itemSubCategoryLink, err := aadiscl.GetItemsSubCategoryLink(itemSubCategoryLinkId)
	if err != nil {
		return adm.ItemsSubCategoryLink{}, err
	}

	return itemSubCategoryLink, nil

}

func (scls *subCategoryLinkService) GetItemsSubCategoryLinkByItemPrivate(currentUser adu.AppliUserLogin, itemId int32) (aasisc.ItemSubCategoryRp, error) {

	itemSubCategoryLink, err := aadiscl.GetItemsSubCategoryLinkByItem(itemId)
	if err != nil {
		return aasisc.ItemSubCategoryRp{}, err
	}

	itemsSubCategoryPublicResponse := aasisc.ItemsSubCategoryResponse{
		IscID:          itemSubCategoryLink.IscID,
		IscName:        itemSubCategoryLink.IscName,
		IscDescription: aus.NullStringToString(itemSubCategoryLink.IscDescription),
		IscPictureUrl:  aus.NullStringToString(itemSubCategoryLink.IscPictureUrl),
	}

	return aasisc.ItemSubCategoryRp(itemsSubCategoryPublicResponse), err

}

func (scls *subCategoryLinkService) InsertItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.CreateItemsSubCategoryLinkParams) (adm.ItemsSubCategoryLink, error) {

	var itemSubCategoryLink adm.ItemsSubCategoryLink
	res, err := aadiscl.InsertItemsSubCategoryLink(itemSubCategoryLinkParams)
	if err != nil {
		return itemSubCategoryLink, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return itemSubCategoryLink, err
	}

	itemSubCategoryLink, err = aadiscl.GetItemsSubCategoryLink(int32(id))
	if err != nil {
		return itemSubCategoryLink, err
	}

	return itemSubCategoryLink, nil
}

func (scls *subCategoryLinkService) UpdateSubItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkParams) (int64, error) {

	rows, err := aadiscl.UpdateItemsSubCategoryLink(itemSubCategoryLinkParams)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (scls *subCategoryLinkService) UpdateSubItemsCategoryByItemLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkParams adm.UpdateItemsSubCategoryLinkByItemParams) (int64, error) {

	rows, err := aadiscl.UpdateItemsSubCategoryByItemLink(itemSubCategoryLinkParams)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (scls *subCategoryLinkService) DeleteItemsSunCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (int64, error) {

	rows, err := aadiscl.DeleteItemsSubCategoryLink(itemSubCategoryLinkId)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (scls *subCategoryLinkService) ListItemsSubCategoryLinkPrivate(currentUser adu.AppliUserLogin) ([]adm.ItemsSubCategoryLink, error) {

	rawItems, err := aadiscl.ListItemsSubCategoryLink()
	if err != nil {
		return nil, err
	}

	return rawItems, nil

}
