package stop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

const ResourceId = "ab75c33d-3a26-4342-b36a-6e5fef0a3ac3"
const ResourceCurrentId = "1c08a38c-ae09-46d2-8926-4f9d25cb0630"

type StopClient struct {
	APIKey string
}

func NewStopClient(key string) *StopClient {
	return &StopClient{APIKey: key}
}

// Get the nominal locations of public transport stops
func (c *StopClient) Get() ([]*Stop, error) {
	return c.getStops(ResourceId)
}

// Get the current locations of the bus stops including temporarily moved ones
func (c *StopClient) GetCurrent() ([]*Stop, error) {
	return c.getStops(ResourceCurrentId)
}

func (c *StopClient) getStops(resourceId string) ([]*Stop, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.stopURL(resourceId).String())
	if err != nil {
		return nil, err
	}

	var response waw.WawResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	return c.extractStops(response), nil
}

func (c *StopClient) newRequest(resourceId string) *StopRequest {
	return &StopRequest{
		ApiKey:     c.APIKey,
		ResourceId: resourceId,
	}
}

func (c *StopClient) stopURL(resourceId string) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/dbstore_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest(resourceId).ToValues().Encode()
	return url
}

func (c *StopClient) extractStops(response waw.WawResponse) []*Stop {
	stops := make([]*Stop, len(response.Result))
	for i, record := range response.Result {
		stops[i] = NewStopFrom(record)
	}
	return stops
}
