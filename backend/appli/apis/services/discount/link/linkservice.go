package link

import (
	aus "github.com/OzarmCtz/e_shop_backend_v1/app/utils/string"

	aaddl "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/discount/link"
	adu "github.com/OzarmCtz/e_shop_backend_v1/appli/apis/domain/users"
	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	DiscountLinkService DiscountLinkServiceInterface = &discountLinkService{}
)

type DiscountLinks []adm.DiscountLink
type DiscountLink adm.DiscountLink
type discountLinkService struct{}

type DiscountLinkServiceInterface interface {
	ListDiscountsLinksPrivate(currentUser adu.AppliUserLogin) ([]DiscountLinkPvResponse, error)
	GetDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkId int32) (DiscountLinkPvResponse, error)
	InsertDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkParams adm.CreateDiscountLinkParams) (DiscountLinkPvResponse, error)
	DeleteDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkId int32) (int64, error)
	UpdateDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkParams adm.UpdateDiscountLinkParams) (int64, error)
}

func (dls *discountLinkService) ListDiscountsLinksPrivate(currentUser adu.AppliUserLogin) ([]DiscountLinkPvResponse, error) {
	discountsLinks, err := aaddl.ListDiscountsLinks()
	if err != nil {
		return nil, err
	}

	privateDiscountLinks := make([]DiscountLinkPvResponse, 0, len(discountsLinks))
	for _, discountLink := range discountsLinks {
		privateDiscountLink := DiscountLinkPvResponse{
			DlID:               discountLink.DlID,
			DlDiscountID:       discountLink.DlDiscountID,
			DlItemsID:          aus.NullInt32ToInt(discountLink.DlItemsID),
			DlItemsSubCategory: aus.NullInt32ToInt(discountLink.DlItemsSubCategory),
			DlItemsCategory:    aus.NullInt32ToInt(discountLink.DlItemsCategory),
		}
		privateDiscountLinks = append(privateDiscountLinks, privateDiscountLink)
	}

	return privateDiscountLinks, err
}

func (dls *discountLinkService) GetDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkId int32) (DiscountLinkPvResponse, error) {
	discountLink, err := aaddl.GetDiscountLink(discountLinkId)
	if err != nil {
		return DiscountLinkPvResponse{}, err
	}

	privateDiscountLink := PrivateDiscountLinkResponse{
		DlID:               discountLink.DlID,
		DlDiscountID:       discountLink.DlDiscountID,
		DlItemsID:          aus.NullInt32ToInt(discountLink.DlItemsID),
		DlItemsSubCategory: aus.NullInt32ToInt(discountLink.DlItemsSubCategory),
		DlItemsCategory:    aus.NullInt32ToInt(discountLink.DlItemsCategory),
	}
	return DiscountLinkPvResponse(privateDiscountLink), err
}

func (dls *discountLinkService) InsertDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkParams adm.CreateDiscountLinkParams) (DiscountLinkPvResponse, error) {
	rows, err := aaddl.InsertDiscountLink(discountLinkParams)
	if err != nil {
		return DiscountLinkPvResponse{}, err
	}

	discountLinkId, err := rows.LastInsertId()
	if err != nil {
		return DiscountLinkPvResponse{}, err
	}

	discountLink, err := aaddl.GetDiscountLink(int32(discountLinkId))
	if err != nil {
		return DiscountLinkPvResponse{}, err

	}

	privateDiscountLink := PrivateDiscountLinkResponse{
		DlID:               discountLink.DlID,
		DlDiscountID:       discountLink.DlDiscountID,
		DlItemsID:          aus.NullInt32ToInt(discountLink.DlItemsID),
		DlItemsSubCategory: aus.NullInt32ToInt(discountLink.DlItemsSubCategory),
		DlItemsCategory:    aus.NullInt32ToInt(discountLink.DlItemsCategory),
	}
	return DiscountLinkPvResponse(privateDiscountLink), err
}

func (dls *discountLinkService) DeleteDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkId int32) (int64, error) {
	rows, err := aaddl.DeleteDiscountLink(discountLinkId)
	if err != nil {
		return 0, err
	}
	return rows, err
}

func (dls *discountLinkService) UpdateDiscountLinkPrivate(currentUser adu.AppliUserLogin, discountLinkParams adm.UpdateDiscountLinkParams) (int64, error) {
	rows, err := aaddl.UpdateDiscountLink(discountLinkParams)
	if err != nil {
		return 0, err
	}
	return rows, err
}
