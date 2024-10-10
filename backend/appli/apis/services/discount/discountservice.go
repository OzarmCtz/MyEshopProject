package discount

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"
	aadas "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/discount"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	DiscountService DiscountServiceInterface = &discountService{}
)

type Discounts []adm.Discount
type Discount adm.Discount
type discountService struct{}

type DiscountServiceInterface interface {
	GetDiscountPrivate(currentUser adu.AppliUserLogin, discountId int32) (DiscountPvResponse, error)
	ListDiscountsPrivate(currentUser adu.AppliUserLogin) ([]PrivateDiscountResponse, error)
	InsertDiscountPrivate(currentUser adu.AppliUserLogin, discountParams adm.CreateDiscountParams) (DiscountPvResponse, error)
	UpdateDiscountPrivate(currentUser adu.AppliUserLogin, discountParams adm.UpdateDiscountParams) (int64, error)
	DeleteDiscountPrivate(currentUser adu.AppliUserLogin, discountId int32) (int64, error)
}

func (ds *discountService) GetDiscountPrivate(currentUser adu.AppliUserLogin, discountId int32) (DiscountPvResponse, error) {
	discountRes, err := aadas.GetDiscount(discountId)
	if err != nil {
		return DiscountPvResponse{}, err
	}

	privateDiscount := PrivateDiscountResponse{
		DID:          discountRes.DID,
		DCode:        discountRes.DCode,
		DDescription: aus.NullStringToString(discountRes.DDescription),
		DStartTime:   aus.NullTimeToTime(discountRes.DStartTime),
		DEndTime:     aus.NullTimeToTime(discountRes.DEndTime),
		DZoneTime:    aus.NullStringToString(discountRes.DZoneTime),
		DIsDisabled:  discountRes.DIsDisabled,
		DPriceType:   discountRes.DPriceType,
		DValue:       discountRes.DValue,
	}

	return DiscountPvResponse(privateDiscount), err
}

func (ds *discountService) ListDiscountsPrivate(currentUser adu.AppliUserLogin) ([]PrivateDiscountResponse, error) {
	rawDiscounts, err := aadas.ListDiscounts()
	if err != nil {
		return nil, err
	}

	privateDiscounts := make([]PrivateDiscountResponse, 0, len(rawDiscounts))
	for _, discount := range rawDiscounts {
		privateDiscount := PrivateDiscountResponse{
			DID:          discount.DID,
			DCode:        discount.DCode,
			DDescription: aus.NullStringToString(discount.DDescription),
			DStartTime:   aus.NullTimeToTime(discount.DStartTime),
			DEndTime:     aus.NullTimeToTime(discount.DEndTime),
			DZoneTime:    aus.NullStringToString(discount.DZoneTime),
			DIsDisabled:  discount.DIsDisabled,
			DPriceType:   discount.DPriceType,
			DValue:       discount.DValue,
		}
		privateDiscounts = append(privateDiscounts, privateDiscount)
	}

	return privateDiscounts, err

}

func (ds *discountService) InsertDiscountPrivate(currentUser adu.AppliUserLogin, discountParams adm.CreateDiscountParams) (DiscountPvResponse, error) {

	res, err := aadas.InsertDiscount(discountParams)
	if err != nil {
		return DiscountPvResponse{}, err
	}

	discountId, err := res.LastInsertId()
	if err != nil {
		return DiscountPvResponse{}, err
	}

	discountRes, err := aadas.GetDiscount(int32(discountId))
	if err != nil {
		return DiscountPvResponse{}, err
	}

	privateDiscount := PrivateDiscountResponse{
		DID:          discountRes.DID,
		DCode:        discountRes.DCode,
		DDescription: aus.NullStringToString(discountRes.DDescription),
		DStartTime:   aus.NullTimeToTime(discountRes.DStartTime),
		DEndTime:     aus.NullTimeToTime(discountRes.DEndTime),
		DZoneTime:    aus.NullStringToString(discountRes.DZoneTime),
		DIsDisabled:  discountRes.DIsDisabled,
		DPriceType:   discountRes.DPriceType,
		DValue:       discountRes.DValue,
	}

	return DiscountPvResponse(privateDiscount), err
}

func (ds *discountService) UpdateDiscountPrivate(currentUser adu.AppliUserLogin, discountParams adm.UpdateDiscountParams) (int64, error) {
	rows, err := aadas.UpdateDiscount(discountParams)
	if err != nil {
		return 0, err
	}
	return rows, err
}

func (ds *discountService) DeleteDiscountPrivate(currentUser adu.AppliUserLogin, discountId int32) (int64, error) {
	rows, err := aadas.DeleteDiscount(discountId)
	if err != nil {
		return 0, err
	}
	return rows, err
}
