package vehicle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	req, err := http.NewRequest(http.MethodPost, c.vehiclesURL(vq).String(), bytes.NewReader([]byte{}))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resBody))

	var response VehiclePositionResponse
	var errorResponse waw.ErrorResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		err = json.Unmarshal(resBody, &errorResponse)
		if err != nil {
			return nil, errors.New("unidentified API error occured")
		} else {
			return nil, &waw.WarsawAPIError{ErrorMessage: string(resBody)}
		}
	}

	return response.Result, nil
}

func (c *VehicleClient) vehiclesURL(vq *VehicleQuery) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/busestrams_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(vq).ToValues().Encode()
	return url
}
