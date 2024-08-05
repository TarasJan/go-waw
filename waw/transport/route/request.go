package route

import (
	"net/url"
)

type RoutesRequest struct {
	ApiKey string `json:"apikey"`
}

func (rr *RoutesRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", rr.ApiKey)

	return values
}
