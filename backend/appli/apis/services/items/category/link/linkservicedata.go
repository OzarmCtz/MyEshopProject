package link

type ItemsCategoryLinkResponse struct {
	IcID          int32  `json:"ic_id"`
	IcName        string `json:"ic_name"`
	IcDescription string `json:"ic_description"`
	IcPictureUrl  string `json:"ic_picture_url"`
}

type ItemCategoryLinkRp ItemsCategoryLinkResponse
