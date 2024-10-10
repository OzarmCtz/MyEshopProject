package category

type ItemsCategoryResponse struct {
	IcID          int32  `json:"ic_id"`
	IcName        string `json:"ic_name"`
	IcDescription string `json:"ic_description"`
	IcPictureUrl  string `json:"ic_picture_url"`
}

type ItemCategoryRp ItemsCategoryResponse

type ListItemsCategoryResponse struct {
	IcID            int32  `json:"ic_id"`
	IcName          string `json:"ic_name"`
	IcDescription   string `json:"ic_description"`
	IcPictureUrl    string `json:"ic_picture_url"`
	IcOnIsc         int64  `json:"ic_on_isc"`
	TotalItemsCount int64  `json:"total_items_count"`
}

type ListItemCategoryRp ListItemsCategoryResponse
