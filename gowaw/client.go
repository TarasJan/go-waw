package gowaw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gowaw/gowaw/vehicle"
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

func (c *Client) Fetch() []vehicle.VehicleLocation {
	urlBase := fmt.Sprintf("%s/api/action/busestrams_get", APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(vehicle.Bus).ToValues().Encode()

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewReader([]byte{}))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response vehicle.VehiclePositionResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
	}

	return response.Result
}
