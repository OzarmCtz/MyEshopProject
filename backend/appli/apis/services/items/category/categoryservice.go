package category

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	adic "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items/category"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	ItemsCategoryService itemsCategoryServiceInterface = &itemsCategoryService{}
)

type ItemsCategory []adm.ItemsCategory
type ItemCategory adm.ItemsCategory
type itemsCategoryService struct{}

type itemsCategoryServiceInterface interface {
	GetItemsCategoryPublic(itemCategoryID int32) (ItemCategoryRp, error)
	ListItemsCategoryPublic() ([]ItemCategoryRp, error)
	InserItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryParams adm.CreateItemCategoryParams) (ItemCategoryRp, error)
	UpdateItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryParams adm.UpdateItemCategoryParams) (int64, error)
	DeleteItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryID int32) (int64, error)
	ListItemsCategoryPrivate(currentUser adu.AppliUserLogin) ([]ListItemCategoryRp, error)
}

func (is *itemsCategoryService) GetItemsCategoryPublic(itemCategoryID int32) (ItemCategoryRp, error) {

	itemCategoryRes, err := adic.GetItemsCategory(itemCategoryID)
	if err != nil {
		return ItemCategoryRp{}, err
	}

	itemsCategoryPublicResponse := ItemsCategoryResponse{
		IcID:          itemCategoryRes.IcID,
		IcName:        itemCategoryRes.IcName,
		IcDescription: aus.NullStringToString(itemCategoryRes.IcDescription),
		IcPictureUrl:  aus.NullStringToString(itemCategoryRes.IcPictureUrl),
	}

	return ItemCategoryRp(itemsCategoryPublicResponse), err

}

func (is *itemsCategoryService) ListItemsCategoryPublic() ([]ItemCategoryRp, error) {

	rawItems, err := adic.ListItemsCategory()
	if err != nil {
		return nil, err
	}

	publicItems := make([]ItemCategoryRp, 0, len(rawItems))
	for _, item := range rawItems {
		publicItem := ItemCategoryRp{
			IcID:          item.IcID,
			IcName:        item.IcName,
			IcDescription: aus.NullStringToString(item.IcDescription),
			IcPictureUrl:  aus.NullStringToString(item.IcPictureUrl),
		}
		publicItems = append(publicItems, publicItem)
	}

	return publicItems, nil
}

func (is *itemsCategoryService) ListItemsCategoryPrivate(currentUser adu.AppliUserLogin) ([]ListItemCategoryRp, error) {

	rawItems, err := adic.ListItemsCategoryAndOccurence()
	if err != nil {
		return nil, err
	}

	publicItems := make([]ListItemCategoryRp, 0, len(rawItems))
	for _, item := range rawItems {
		publicItem := ListItemCategoryRp{
			IcID:            item.IcID,
			IcName:          item.IcName,
			IcDescription:   aus.NullStringToString(item.IcDescription),
			IcPictureUrl:    aus.NullStringToString(item.IcPictureUrl),
			IcOnIsc:         item.IcOnIsc,
			TotalItemsCount: item.TotalItemsCount,
		}
		publicItems = append(publicItems, publicItem)
	}

	return publicItems, nil
}

func (is *itemsCategoryService) InserItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryParams adm.CreateItemCategoryParams) (ItemCategoryRp, error) {

	res, err := adic.InsertItemCategory(itemCategoryParams)
	if err != nil {
		return ItemCategoryRp{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return ItemCategoryRp{}, err
	}

	itemCategory, err := adic.GetItemsCategory(int32(id))
	if err != nil {
		return ItemCategoryRp{}, err
	}

	itemCategoryPublicResponse := ItemsCategoryResponse{
		IcID:          itemCategory.IcID,
		IcName:        itemCategory.IcName,
		IcDescription: aus.NullStringToString(itemCategory.IcDescription),
		IcPictureUrl:  aus.NullStringToString(itemCategory.IcPictureUrl),
	}

	return ItemCategoryRp(itemCategoryPublicResponse), nil

}

func (is *itemsCategoryService) UpdateItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryParams adm.UpdateItemCategoryParams) (int64, error) {

	rows, err := adic.UpdateItemCategory(itemCategoryParams)
	if err != nil {
		return rows, err
	}
	return rows, err

}

func (is *itemsCategoryService) DeleteItemsCategoryPrivate(currentUser adu.AppliUserLogin, itemCategoryID int32) (int64, error) {

	rows, err := adic.DeleteItemCategory(itemCategoryID)
	if err != nil {
		return rows, err
	}

	return rows, nil

}
