package wishlist

import "time"

type PrivateWishListResponse struct {
	WlID         int32     `json:"wl_id"`
	WlUserID     int32     `json:"wl_user_id"`
	WlItemsID    int32     `json:"wl_items_id"`
	WlTimesAdded time.Time `json:"wl_times_added"`
}

type WishListPvResponse PrivateWishListResponse
