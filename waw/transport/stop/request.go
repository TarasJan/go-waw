package stop

import (
	"net/url"
)

type StopRequest struct {
	ResourceId string `json:"id"`
	ApiKey     string `json:"apikey"`
}

func (sr *StopRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", sr.ApiKey)
	values.Set("id", sr.ResourceId)

	return values
}
