package dictionary

import (
	"net/url"
)

type DictionaryRequest struct {
	ResourceId string `json:"resource_id"`
	ApiKey     string `json:"apikey"`
}

func (dr *DictionaryRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", dr.ApiKey)
	values.Set("resource_id", dr.ResourceId)

	return values
}
