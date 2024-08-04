package timetable

import (
	"net/url"
)

type TimetableRequest struct {
	ResourceId string `json:"resource_id"`
	ApiKey     string `json:"apikey"`
}

func (tr *TimetableRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", tr.ApiKey)
	values.Set("resource_id", tr.ResourceId)

	return values
}
