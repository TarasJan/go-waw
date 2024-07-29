package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gowaw/gowaw"
	"gowaw/gowaw/transport/vehicle"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{APIKey: key}
}

func (c *Client) newRequest(vt vehicle.VehicleType) *vehicle.VehiclePositionRequest {
	return &vehicle.VehiclePositionRequest{
		ApiKey:     c.APIKey,
		ResourceId: vehicle.VehicleRequestResourcId,
		Type:       vt,
	}
}

func (c *Client) FetchBuses() ([]vehicle.VehicleLocation, error) {
	return c.fetch(vehicle.Bus)
}

func (c *Client) FetchTrams() ([]vehicle.VehicleLocation, error) {
	return c.fetch(vehicle.Tram)
}

func (c *Client) fetch(vt vehicle.VehicleType) ([]vehicle.VehicleLocation, error) {
	req, err := http.NewRequest(http.MethodPost, c.vehiclesURL(vt).String(), bytes.NewReader([]byte{}))
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

	var response vehicle.VehiclePositionResponse
	var errorResponse gowaw.ErrorResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		err = json.Unmarshal(resBody, &errorResponse)
		if err != nil {
			return nil, errors.New("unidentified API error occured")
		} else {
			return nil, &gowaw.UnauthorizedAccessError{}
		}
	}

	return response.Result, nil
}

func (c *Client) vehiclesURL(vt vehicle.VehicleType) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/busestrams_get", gowaw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(vt).ToValues().Encode()
	return url
}
