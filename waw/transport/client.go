package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
	"github.com/TarasJan/go-waw/waw/transport/vehicle"
)

type Client struct {
	APIKey string
}

func NewClient(key string) *Client {
	return &Client{APIKey: key}
}

func (c *Client) newRequest(vq *vehicle.VehicleQuery) *vehicle.VehiclePositionRequest {
	return &vehicle.VehiclePositionRequest{
		ApiKey:        c.APIKey,
		ResourceId:    vehicle.VehicleRequestResourcId,
		Type:          vq.Type,
		Line:          vq.Line,
		BrigadeNumber: vq.BrigadeNumber,
	}
}

func (c *Client) FetchBuses(options ...vehicle.VehicleQueryOption) ([]vehicle.VehicleLocation, error) {
	return c.fetch(vehicle.NewBusQuery(options...))
}

func (c *Client) FetchTrams(options ...vehicle.VehicleQueryOption) ([]vehicle.VehicleLocation, error) {
	return c.fetch(vehicle.NewTramQuery(options...))
}

func (c *Client) fetch(vq *vehicle.VehicleQuery) ([]vehicle.VehicleLocation, error) {
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

	var response vehicle.VehiclePositionResponse
	var errorResponse waw.ErrorResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		err = json.Unmarshal(resBody, &errorResponse)
		if err != nil {
			return nil, errors.New("unidentified API error occured")
		} else {
			return nil, &waw.UnauthorizedAccessError{}
		}
	}

	return response.Result, nil
}

func (c *Client) vehiclesURL(vq *vehicle.VehicleQuery) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/busestrams_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(vq).ToValues().Encode()
	return url
}
