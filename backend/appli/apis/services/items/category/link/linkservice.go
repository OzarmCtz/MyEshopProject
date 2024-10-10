package link

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	adicl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/category/link"
	aadisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/sub/category"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aasic "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items/category"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	CategoryLinkService categoryLinkServiceInterface = &categoryLinkService{}
)

type ItemsCategoryLink []adm.ItemsCategoryLink
type ItemCategoryLink adm.ItemsCategoryLink
type categoryLinkService struct{}

type categoryLinkServiceInterface interface {
	ListItemsCategoryByCategoryPublic(categoryId int32) ([]adm.ItemsSubCategory, error)
	ListItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin) ([]adm.ItemsCategoryLink, error)
	GetItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkId int32) (adm.ItemsCategoryLink, error)
	InsertItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkParams adm.CreateItemsCategoryLinkParams) (adm.ItemsCategoryLink, error)
	UpdateItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkParams adm.UpdateItemsCategoryLinkParams) (int64, error)
	DeleteItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkId int32) (int64, error)
	GetItemsCategoryLinkBySubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (aasic.ItemCategoryRp, error)
	UpdateItemsCategoryLinkBySubCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkBySubCategoryParams adm.UpdateItemsCategoryLinkBySubCategoryParams) (int64, error)
}

// ListItemsCategoryBySubCategory is a method to list sub category by category
func (cls *categoryLinkService) ListItemsCategoryByCategoryPublic(categoryId int32) ([]adm.ItemsSubCategory, error) {

	rawItems, err := adicl.ListItemsCategoryLinkByCategory(categoryId)
	if err != nil {
		return nil, err
	}

	publicItems := make([]adm.ItemsSubCategory, 0, len(rawItems))

	for _, item := range rawItems {
		publicItem, err := aadisc.GetItemsSubCategory(item.IclItemsSubCategoryID)
		if err != nil {
			return nil, err
		}
		publicItems = append(publicItems, publicItem)

	}

	return publicItems, nil

}

func (cls *categoryLinkService) GetItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkId int32) (adm.ItemsCategoryLink, error) {

	itemCategoryLink, err := adicl.GetItemsCategoryLink(itemCategoryLinkId)
	if err != nil {
		return adm.ItemsCategoryLink{}, err
	}

	return itemCategoryLink, nil

}

func (cls *categoryLinkService) GetItemsCategoryLinkBySubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryLinkId int32) (aasic.ItemCategoryRp, error) {

	itemCategoryLink, err := adicl.GetItemCategoryLinkBySubCategory(itemSubCategoryLinkId)
	if err != nil {
		return aasic.ItemCategoryRp{}, err
	}

	itemsCategoryResponse := aasic.ItemsCategoryResponse{
		IcID:          itemCategoryLink.IcID,
		IcName:        itemCategoryLink.IcName,
		IcDescription: aus.NullStringToString(itemCategoryLink.IcDescription),
		IcPictureUrl:  aus.NullStringToString(itemCategoryLink.IcPictureUrl),
	}

	return aasic.ItemCategoryRp(itemsCategoryResponse), nil

}

func (cls *categoryLinkService) InsertItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkParams adm.CreateItemsCategoryLinkParams) (adm.ItemsCategoryLink, error) {

	var itemCategoryLink adm.ItemsCategoryLink
	res, err := adicl.InsertItemsCategoryLink(itemCategoryLinkParams)
	if err != nil {
		return itemCategoryLink, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return itemCategoryLink, err
	}

	itemCategoryLink, err = adicl.GetItemsCategoryLink(int32(id))
	if err != nil {
		return itemCategoryLink, err
	}

	return itemCategoryLink, nil
}

func (cls *categoryLinkService) UpdateItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkParams adm.UpdateItemsCategoryLinkParams) (int64, error) {

	rows, err := adicl.UpdateItemsCategoryLink(itemCategoryLinkParams)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (cls *categoryLinkService) UpdateItemsCategoryLinkBySubCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkBySubCategoryParams adm.UpdateItemsCategoryLinkBySubCategoryParams) (int64, error) {

	rows, err := adicl.UpdateItemsCategoryLinkBySubCategory(itemCategoryLinkBySubCategoryParams)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (cls *categoryLinkService) DeleteItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin, itemCategoryLinkId int32) (int64, error) {

	rows, err := adicl.DeleteItemsCategoryLink(itemCategoryLinkId)
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (cls *categoryLinkService) ListItemsCategoryLinkPrivate(currentUser adu.AppliUserLogin) ([]adm.ItemsCategoryLink, error) {

	rawItems, err := adicl.ListItemsCategoryLink()
	if err != nil {
		return nil, err
	}

	return rawItems, nil

}
