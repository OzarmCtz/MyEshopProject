package itemservice

import (
	"errors"
	"time"

	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	adi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/items"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
	aua "github.com/OzarmCtz/e_shop_backend_v1/appli/utils/auth"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type Items []adm.User
type Item adm.Item
type itemsService struct{}

type itemsServiceInterface interface {
	InsertItems(currentUser adu.AppliUserLogin, itemParams adm.CreateItemParams) (ItemPvResponse, error)
	GetItemsPublic(itemID int32) (ItemPbResponse, error)
	ListItemsPublic() ([]PublicItemResponse, error)
	ListItemsPrivate(currentUser adu.AppliUserLogin) ([]adm.ListItemsRow, error)
	UpdateItems(currentUser adu.AppliUserLogin, itemParams adm.UpdateItemParams) (int64, error)
	DeleteItems(currentUser adu.AppliUserLogin, itemID int32) (int64, error)
	GetItemsPrivate(currentUser adu.AppliUserLogin, itemID int32) (ItemPvResponse, error)
}

// PUBLIC

func (is *itemsService) GetItemsPublic(itemID int32) (ItemPbResponse, error) {

	itemRes, err := adi.GetItems(itemID)
	if err != nil {
		return ItemPbResponse{}, err
	}

	if itemRes.IIsDisabled {
		return ItemPbResponse{}, errors.New("Item is disabled and cannot be accessed")
	}

	publicItemResponse := PublicItemResponse{
		IID:          itemRes.IID,
		ITitle:       itemRes.ITitle,
		IDescription: aus.NullStringToString(itemRes.IDescription),
		IPrice:       itemRes.IPrice,
		IQuantity:    aus.NullInt32ToInt(itemRes.IQuantity),
		IPictureUrl:  aus.NullStringToString(itemRes.IPictureUrl),
		IReleaseDate: itemRes.IReleaseDate,
	}

	return ItemPbResponse(publicItemResponse), err
}

func (is *itemsService) ListItemsPublic() ([]PublicItemResponse, error) {

	rawItems, err := adi.ListItemByActivity(false)
	if err != nil {
		return nil, err
	}

	publicItems := make([]PublicItemResponse, 0, len(rawItems))
	for _, item := range rawItems {
		publicItem := PublicItemResponse{
			IID:          item.IID,
			ITitle:       item.ITitle,
			IDescription: aus.NullStringToString(item.IDescription),
			IPrice:       item.IPrice,
			IQuantity:    aus.NullInt32ToInt(item.IQuantity),
			IPictureUrl:  aus.NullStringToString(item.IPictureUrl),
			IReleaseDate: item.IReleaseDate,
		}
		publicItems = append(publicItems, publicItem)
	}

	return publicItems, nil
}

// PRIVATE FUNCTIONS
func (is *itemsService) ListItemsPrivate(currentUser adu.AppliUserLogin) ([]adm.ListItemsRow, error) {

	userIsAdmin, err := aua.IsRealyAdmin(currentUser.User.UID)
	if err != nil {
		return nil, err
	}

	userIsSuperAdmin, err := aua.IsRealySuperAdmin(currentUser.User.UID)
	if err != nil {
		return nil, err
	}

	if !userIsAdmin && !userIsSuperAdmin {
		return nil, errors.New("you are not authorized to view this resource")
	}

	items, err := adi.ListItems()

	if err != nil {
		return items, err
	}

	return items, err
}

func (is *itemsService) GetItemsPrivate(currentUser adu.AppliUserLogin, itemID int32) (ItemPvResponse, error) {

	userIsAdmin, err := aua.IsRealyAdmin(currentUser.User.UID)
	if err != nil {
		return ItemPvResponse{}, err
	}

	userIsSuperAdmin, err := aua.IsRealySuperAdmin(currentUser.User.UID)
	if err != nil {
		return ItemPvResponse{}, err
	}

	if !userIsAdmin && !userIsSuperAdmin {
		return ItemPvResponse{}, errors.New("you are not authorized to view this resource")
	}

	itemRes, err := adi.GetItems(itemID)
	if err != nil {
		return ItemPvResponse{}, err
	}

	privateItemResponse := PrivateItemResponse{
		IID:          itemRes.IID,
		ITitle:       itemRes.ITitle,
		IDescription: aus.NullStringToString(itemRes.IDescription),
		IPrice:       itemRes.IPrice,
		IQuantity:    aus.NullInt32ToInt(itemRes.IQuantity),
		IFilePath:    aus.NullStringToString(itemRes.IFilePath),
		IIsDisabled:  itemRes.IIsDisabled,
		IPictureUrl:  aus.NullStringToString(itemRes.IPictureUrl),
		IReleaseDate: itemRes.IReleaseDate,
	}

	return ItemPvResponse(privateItemResponse), err
}

func (is *itemsService) InsertItems(currentUser adu.AppliUserLogin, itemParams adm.CreateItemParams) (ItemPvResponse, error) {

	itemParams.IReleaseDate = time.Now()

	res, err := adi.InsertItem(itemParams)
	if err != nil {
		return ItemPvResponse{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return ItemPvResponse{}, err
	}

	itemRes, err := adi.GetItems(int32(id))
	if err != nil {
		return ItemPvResponse{}, err
	}

	privateItemResponse := PrivateItemResponse{
		IID:          itemRes.IID,
		ITitle:       itemRes.ITitle,
		IDescription: aus.NullStringToString(itemRes.IDescription),
		IPrice:       itemRes.IPrice,
		IQuantity:    aus.NullInt32ToInt(itemRes.IQuantity),
		IFilePath:    aus.NullStringToString(itemRes.IFilePath),
		IIsDisabled:  itemRes.IIsDisabled,
		IPictureUrl:  aus.NullStringToString(itemRes.IPictureUrl),
		IReleaseDate: itemRes.IReleaseDate,
	}

	return ItemPvResponse(privateItemResponse), nil

}

func (is *itemsService) UpdateItems(currentUser adu.AppliUserLogin, itemParams adm.UpdateItemParams) (int64, error) {

	itemParams.IReleaseDate = time.Now()

	rows, err := adi.UpdateItem(itemParams)
	if err != nil {
		return rows, err
	}

	return rows, err
}

func (is *itemsService) DeleteItems(currentUser adu.AppliUserLogin, itemID int32) (int64, error) {
	rows, err := adi.DeleteItem(itemID)
	if err != nil {
		return rows, err
	}

	return rows, err
}
