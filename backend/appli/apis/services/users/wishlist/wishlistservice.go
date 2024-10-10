package wishlist

import (
	"errors"
	"time"

	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	aduw "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users/wishlist"
	aasi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	WishListService wishListServiceInterface = &wishListService{}
)

type WishLists []adm.UsersWishlist
type WishList adm.UsersWishlist
type wishListService struct{}

type wishListServiceInterface interface {
	GetUserWishListByUserPrivate(currentUser adu.AppliUserLogin, wishListUserId int32) ([]aasi.ItemPbResponse, error)
	InsertUserWishListPrivate(currentUser adu.AppliUserLogin, wishListParams adm.CreateUserWishListParams) (PrivateWishListResponse, error)
	DeleteUserWishListByUserAndItemPrivate(currentUser adu.AppliUserLogin, wishListParams adm.DeleteUserWishListByUserAndItemsParams) (int64, error)
}

func (wls *wishListService) GetUserWishListByUserPrivate(currentUser adu.AppliUserLogin, wishListUserId int32) ([]aasi.ItemPbResponse, error) {

	var items []aasi.ItemPbResponse
	if wishListUserId == currentUser.User.UID {

		wishList, err := aduw.ListUserWishListByUser(wishListUserId)

		if err != nil {
			return items, err
		}

		userItems := make([]aasi.ItemPbResponse, 0, len(wishList))

		for _, item := range wishList {
			useItem, err := aasi.ItemsService.GetItemsPublic(item.WlItemsID.Int32)
			if err != nil {
				return items, err
			}
			userItems = append(userItems, useItem)
		}

		return userItems, nil

	}

	return items, errors.New("user does not have permission to view this wishlist")
}

func (wls *wishListService) InsertUserWishListPrivate(currentUser adu.AppliUserLogin, wishListParams adm.CreateUserWishListParams) (PrivateWishListResponse, error) {

	if wishListParams.WlUserID == currentUser.User.UID {

		wishListParams.WlTimesAdded = time.Now()

		res, err := aduw.InsertUserWishList(wishListParams)
		if err != nil {
			return PrivateWishListResponse{}, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return PrivateWishListResponse{}, err
		}

		wishList, err := aduw.GetUserWishList(int32(id))
		if err != nil {
			return PrivateWishListResponse{}, err
		}

		privateWishListResponse := PrivateWishListResponse{
			WlID:         wishList.WlID,
			WlUserID:     wishList.WlUserID,
			WlItemsID:    aus.NullInt32ToInt(wishList.WlItemsID),
			WlTimesAdded: wishList.WlTimesAdded,
		}

		return PrivateWishListResponse(privateWishListResponse), err

	}

	return PrivateWishListResponse{}, errors.New("user does not have permission to add to this wishlist")
}

func (wls *wishListService) DeleteUserWishListByUserAndItemPrivate(currentUser adu.AppliUserLogin, wishListParams adm.DeleteUserWishListByUserAndItemsParams) (int64, error) {

	if wishListParams.WlUserID == currentUser.User.UID {

		rows, err := aduw.DeleteUserWishListByUserAndItem(wishListParams)
		if err != nil {
			return rows, err
		}

		return rows, nil
	}
	return 0, errors.New("user does not have permission to delete this wishlist")
}
