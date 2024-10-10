package itemservice

import "time"

type PublicItemResponse struct {
	IID          int32     `json:"i_id"`
	ITitle       string    `json:"i_title"`
	IDescription string    `json:"i_description"`
	IPrice       string    `json:"i_price"`
	IQuantity    int32     `json:"i_quantity"`
	IPictureUrl  string    `json:"i_picture_url"`
	IReleaseDate time.Time `json:"i_release_date"`
}

type PrivateItemResponse struct {
	IID          int32     `json:"i_id"`
	ITitle       string    `json:"i_title"`
	IDescription string    `json:"i_description"`
	IPrice       string    `json:"i_price"`
	IQuantity    int32     `json:"i_quantity"`
	IPictureUrl  string    `json:"i_picture_url"`
	IFilePath    string    `json:"i_file_path"`
	IIsDisabled  bool      `json:"i_is_disabled"`
	IReleaseDate time.Time `json:"i_release_date"`
}
type ItemPbResponse PublicItemResponse
type ItemPvResponse PrivateItemResponse
