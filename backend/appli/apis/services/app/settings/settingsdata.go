package settings

import "time"

type PublicAppSettingResponse struct {
	AsKey   string `json:"as_key"`
	AsValue string `json:"as_value"`
}

type PrivateAppSettingResponse struct {
	AsID          int32     `json:"as_id"`
	AsKey         string    `json:"as_key"`
	AsValue       string    `json:"as_value"`
	AsDescription string    `json:"as_description"`
	AsLastUpdated time.Time `json:"as_last_updated"`
}

type AppSettingPbResponse PublicAppSettingResponse
type AppSettingPvResponse PrivateAppSettingResponse
