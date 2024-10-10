package category

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	adisc "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/sub/category"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	ItemsSubCategoryService itemsSubCategoryServiceInterface = &itemsSubCategoryService{}
)

type ItemsSubCategory []adm.ItemsSubCategory
type ItemSubCategory adm.ItemsSubCategory
type itemsSubCategoryService struct{}

type itemsSubCategoryServiceInterface interface {
	GetItemsSubCategoryPublic(itemSubCategoryID int32) (ItemSubCategoryRp, error)
	ListItemsSubCategoryPublic() ([]ItemSubCategoryRp, error)
	InserItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryParams adm.CreateItemSubCategoryParams) (ItemSubCategoryRp, error)
	UpdateItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryParams adm.UpdateItemSubCategoryParams) (int64, error)
	DeleteItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryID int32) (int64, error)
	GetItemsSubCategoryAndCategoryLinkedPrivate(currentUser adu.AppliUserLogin, itemSubCategoryID int32) (ItemSubCategoryAndCategoryLinkedRp, error)
	ListItemsSubCategoryAndCategoryLinkedPrivate(currentUser adu.AppliUserLogin) ([]ItemSubCategoryAndCategoryLinkedRp, error)
}

func (is *itemsSubCategoryService) GetItemsSubCategoryAndCategoryLinkedPrivate(currentUser adu.AppliUserLogin, itemSubCategoryID int32) (ItemSubCategoryAndCategoryLinkedRp, error) {

	itemSubCategoryRes, err := adisc.GetItemsSubCategoryAndCategoryLinked(itemSubCategoryID)
	if err != nil {
		return ItemSubCategoryAndCategoryLinkedRp{}, err
	}

	itemsSubCategoryPublicResponse := ItemsSubCategoryAndCategoryLinkedResponse{
		IscID:          itemSubCategoryRes.IscID,
		IscName:        itemSubCategoryRes.IscName,
		IscDescription: aus.NullStringToString(itemSubCategoryRes.IscDescription),
		IscPictureUrl:  aus.NullStringToString(itemSubCategoryRes.IscPictureUrl),
		IcName:         aus.NullStringToString(itemSubCategoryRes.IcName),
	}

	return ItemSubCategoryAndCategoryLinkedRp(itemsSubCategoryPublicResponse), err

}

func (is *itemsSubCategoryService) GetItemsSubCategoryPublic(itemSubCategoryID int32) (ItemSubCategoryRp, error) {

	itemSubCategoryRes, err := adisc.GetItemsSubCategory(itemSubCategoryID)
	if err != nil {
		return ItemSubCategoryRp{}, err
	}

	itemsSubCategoryPublicResponse := ItemsSubCategoryResponse{
		IscID:          itemSubCategoryRes.IscID,
		IscName:        itemSubCategoryRes.IscName,
		IscDescription: aus.NullStringToString(itemSubCategoryRes.IscDescription),
		IscPictureUrl:  aus.NullStringToString(itemSubCategoryRes.IscPictureUrl),
	}

	return ItemSubCategoryRp(itemsSubCategoryPublicResponse), err

}

func (is *itemsSubCategoryService) ListItemsSubCategoryAndCategoryLinkedPrivate(currentUser adu.AppliUserLogin) ([]ItemSubCategoryAndCategoryLinkedRp, error) {

	rawItems, err := adisc.ListItemsSubCategoryAndCategoryLinked()
	if err != nil {
		return nil, err
	}

	privateSubItems := make([]ItemSubCategoryAndCategoryLinkedRp, 0, len(rawItems))
	for _, item := range rawItems {
		privateSubItem := ItemsSubCategoryAndCategoryLinkedResponse{
			IscID:          item.IscID,
			IscName:        item.IscName,
			IscDescription: aus.NullStringToString(item.IscDescription),
			IscPictureUrl:  aus.NullStringToString(item.IscPictureUrl),
			IcName:         aus.NullStringToString(item.IcName),
			ItemCount:      item.ItemCount,
		}
		privateSubItems = append(privateSubItems, ItemSubCategoryAndCategoryLinkedRp(privateSubItem))
	}

	return privateSubItems, nil
}

func (is *itemsSubCategoryService) ListItemsSubCategoryPublic() ([]ItemSubCategoryRp, error) {

	rawItems, err := adisc.ListItemsSubCategory()
	if err != nil {
		return nil, err
	}

	publicSubItems := make([]ItemSubCategoryRp, 0, len(rawItems))
	for _, item := range rawItems {
		publicSubItem := ItemSubCategoryRp{
			IscID:          item.IscID,
			IscName:        item.IscName,
			IscDescription: aus.NullStringToString(item.IscDescription),
			IscPictureUrl:  aus.NullStringToString(item.IscPictureUrl),
		}
		publicSubItems = append(publicSubItems, publicSubItem)
	}

	return publicSubItems, nil
}

func (is *itemsSubCategoryService) InserItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryParams adm.CreateItemSubCategoryParams) (ItemSubCategoryRp, error) {

	res, err := adisc.InsertItemSubCategory(itemSubCategoryParams)
	if err != nil {
		return ItemSubCategoryRp{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return ItemSubCategoryRp{}, err
	}

	itemSubCategory, err := adisc.GetItemsSubCategory(int32(id))
	if err != nil {
		return ItemSubCategoryRp{}, err
	}

	itemSubCategoryPublicResponse := ItemsSubCategoryResponse{
		IscID:          itemSubCategory.IscID,
		IscName:        itemSubCategory.IscName,
		IscDescription: aus.NullStringToString(itemSubCategory.IscDescription),
		IscPictureUrl:  aus.NullStringToString(itemSubCategory.IscPictureUrl),
	}

	return ItemSubCategoryRp(itemSubCategoryPublicResponse), nil

}

func (is *itemsSubCategoryService) UpdateItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryParams adm.UpdateItemSubCategoryParams) (int64, error) {

	rows, err := adisc.UpdateItemSubCategory(itemSubCategoryParams)
	if err != nil {
		return rows, err
	}
	return rows, err

}

func (is *itemsSubCategoryService) DeleteItemsSubCategoryPrivate(currentUser adu.AppliUserLogin, itemSubCategoryID int32) (int64, error) {

	rows, err := adisc.DeleteItemSubCategory(itemSubCategoryID)
	if err != nil {
		return rows, err
	}

	return rows, nil

}
