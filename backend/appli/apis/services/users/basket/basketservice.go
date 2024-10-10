package basket

import (
	"errors"
	"time"

	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adub "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users/basket"
	aasi "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/services/items"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var BasketService basketServiceInterface = &basketService{}

type Baskets []adm.UsersBasket
type Basket adm.UsersBasket
type basketService struct{}

type basketServiceInterface interface {
	GetUserBasketByUserPrivate(currentUser adu.AppliUserLogin, basketUserId int32) ([]aasi.ItemPbResponse, error)
	InsertUserBasketPrivate(currentUser adu.AppliUserLogin, basketParams adm.CreateUserBasketParams) (adm.UsersBasket, error)
	DeleteUserBasketByUserAndItemPrivate(currentUser adu.AppliUserLogin, basketParams adm.DeleteUserBasketByUserAndItemsParams) (int64, error)
}

func (bs *basketService) GetUserBasketByUserPrivate(currentUser adu.AppliUserLogin, basketUserId int32) ([]aasi.ItemPbResponse, error) {

	var items []aasi.ItemPbResponse
	if basketUserId == currentUser.User.UID {

		wishList, err := adub.ListUserBasketByUser(basketUserId)

		if err != nil {
			return items, err
		}

		userItems := make([]aasi.ItemPbResponse, 0, len(wishList))

		for _, item := range wishList {
			useItem, err := aasi.ItemsService.GetItemsPublic(item.UbItemsID)
			if err != nil {
				return items, err
			}
			userItems = append(userItems, useItem)
		}

		return userItems, nil

	}

	return items, errors.New("user does not have permission to view this basket")
}

func (bs *basketService) InsertUserBasketPrivate(currentUser adu.AppliUserLogin, basketParams adm.CreateUserBasketParams) (adm.UsersBasket, error) {

	var basket adm.UsersBasket
	if basketParams.UbUserID == currentUser.User.UID {

		basketParams.UbTimeAdded = time.Now()

		res, err := adub.InsertBasketList(basketParams)
		if err != nil {
			return basket, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return basket, err
		}

		basket, err := adub.GetUserBasket(int32(id))
		if err != nil {
			return basket, err
		}

		return basket, err

	}

	return basket, errors.New("user does not have permission to add items to this basket")
}

func (bs *basketService) DeleteUserBasketByUserAndItemPrivate(currentUser adu.AppliUserLogin, basketParams adm.DeleteUserBasketByUserAndItemsParams) (int64, error) {

	if basketParams.UbUserID == currentUser.User.UID {

		rows, err := adub.DeleteUserBasketByUserAndItem(basketParams)
		if err != nil {
			return rows, err
		}

		return rows, nil
	}
	return 0, errors.New("user does not have permission to delete this basket")
}
