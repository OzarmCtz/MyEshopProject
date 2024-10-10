package link

type PrivateDiscountLinkResponse struct {
	DlID               int32 `json:"dl_id"`
	DlDiscountID       int32 `json:"dl_discount_id"`
	DlItemsID          int32 `json:"dl_items_id"`
	DlItemsSubCategory int32 `json:"dl_items_sub_category"`
	DlItemsCategory    int32 `json:"dl_items_category"`
}

type DiscountLinkPvResponse PrivateDiscountLinkResponse
