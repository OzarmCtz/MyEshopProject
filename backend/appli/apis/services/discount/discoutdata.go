package discount

import "time"

type PublicDiscountResponse struct {
	DCode      string `json:"d_code"`
	DPriceType string `json:"d_price_type"`
	DValue     int32  `json:"d_value"`
}

type PrivateDiscountResponse struct {
	DID          int32     `json:"d_id"`
	DCode        string    `json:"d_code"`
	DDescription string    `json:"d_description"`
	DStartTime   time.Time `json:"d_start_time"`
	DEndTime     time.Time `json:"d_end_time"`
	DZoneTime    string    `json:"d_zone_time"`
	DIsDisabled  bool      `json:"d_is_disabled"`
	DPriceType   string    `json:"d_price_type"`
	DValue       int32     `json:"d_value"`
}

type DiscountPvResponse PrivateDiscountResponse
type DiscountPbResponse PublicDiscountResponse
