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
	Line          string      `json:"line"`    // string due to some lines starting with letters N or L
	BrigadeNumber string      `json:"brigade"` // string due to some brigades starting with 0
}

func (vpr *VehiclePositionRequest) ToValues() url.Values {
	values := url.Values{}
	values.Set("apikey", vpr.ApiKey)
	values.Set("resource_id", vpr.ResourceId)
	values.Set("type", fmt.Sprintf("%d", vpr.Type))
	if vpr.Line != "" {
		values.Set("line", vpr.Line)
	}
	if vpr.BrigadeNumber != "" {
		values.Set("brigade", vpr.Line)
	}

	return values
}
