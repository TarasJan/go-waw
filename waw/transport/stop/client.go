package stop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

const ResourceId = "ab75c33d-3a26-4342-b36a-6e5fef0a3ac3"

type StopClient struct {
	APIKey string
}

func NewStopClient(key string) *StopClient {
	return &StopClient{APIKey: key}
}

func (c *StopClient) Get() ([]*Stop, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.stopURL().String())
	if err != nil {
		return nil, err
	}

	var response waw.WawResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	stops := make([]*Stop, len(response.Result))
	for i, record := range response.Result {
		stops[i] = NewStopFrom(record)
	}

	return stops, nil
}

func (c *StopClient) newRequest() *StopRequest {
	return &StopRequest{
		ApiKey:     c.APIKey,
		ResourceId: ResourceId,
	}
}

func (c *StopClient) stopURL() *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/dbstore_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest().ToValues().Encode()
	return url
}
