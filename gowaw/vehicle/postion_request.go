package vehicle

import (
	"fmt"
	"net/url"
)

type VehiclePositionResponse struct {
	Result []VehicleLocation `json:"result"`
}

type VehiclePositionRequest struct {
	ResourceId    string      `json:"resource_id"`
	ApiKey        string      `json:"apikey"`
	Type          VehicleType `json:"type"`
	Line          int64       `json:"line"`
	BrigadeNumber int64       `json:"brigade"`
}

func (vpr *VehiclePositionRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", vpr.ApiKey)
	values.Set("resource_id", vpr.ResourceId)
	values.Set("type", fmt.Sprintf("%d", vpr.Type))
	if vpr.Line > 0 {
		values.Set("line", fmt.Sprintf("%d", vpr.Line))
	}
	if vpr.BrigadeNumber > 0 {
		values.Set("brigade", fmt.Sprintf("%d", vpr.BrigadeNumber))
	}

	return values
}
