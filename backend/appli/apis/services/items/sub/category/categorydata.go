package category

type ItemsSubCategoryResponse struct {
	IscID          int32  `json:"isc_id"`
	IscName        string `json:"isc_name"`
	IscDescription string `json:"isc_description"`
	IscPictureUrl  string `json:"isc_picture_url"`
}

type ItemSubCategoryRp ItemsSubCategoryResponse

type ItemsSubCategoryAndCategoryLinkedResponse struct {
	IscID          int32  `json:"isc_id"`
	IscName        string `json:"isc_name"`
	IscDescription string `json:"isc_description"`
	IscPictureUrl  string `json:"isc_picture_url"`
	IcName         string `json:"ic_name"`
	ItemCount      int64  `json:"item_count"`
}

type ItemSubCategoryAndCategoryLinkedRp ItemsSubCategoryAndCategoryLinkedResponse
