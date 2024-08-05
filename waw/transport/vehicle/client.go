package vehicle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

type VehicleClient struct {
	APIKey string
}

func NewVehicleClient(key string) *VehicleClient {
	return &VehicleClient{APIKey: key}
}

func (c *VehicleClient) newRequest(vq *VehicleQuery) *VehiclePositionRequest {
	return &VehiclePositionRequest{
		ApiKey:        c.APIKey,
		ResourceId:    ResourceId,
		Type:          vq.Type,
		Line:          vq.Line,
		BrigadeNumber: vq.BrigadeNumber,
	}
}

func (c *VehicleClient) FetchBuses(options ...VehicleQueryOption) ([]VehicleLocation, error) {
	return c.fetch(NewBusQuery(options...))
}

func (c *VehicleClient) FetchTrams(options ...VehicleQueryOption) ([]VehicleLocation, error) {
	return c.fetch(NewTramQuery(options...))
}

func (c *VehicleClient) fetch(vq *VehicleQuery) ([]VehicleLocation, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodPost, c.vehiclesURL(vq).String())
	if err != nil {
		return nil, err
	}

	var response VehiclePositionResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	return response.Result, nil
}

func (c *VehicleClient) vehiclesURL(vq *VehicleQuery) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/busestrams_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(vq).ToValues().Encode()
	return url
}
