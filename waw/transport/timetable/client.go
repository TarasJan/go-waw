package timetable

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

const ResourceStopTimetableId = "8b27f4c17-5c50-4a5b-89dd-236b282bc499"
const ResourceLineTimetableId = "88cd555f-6f31-43ca-9de4-66c479ad5942"
const ResourceTimetableId = "e923fa0e-d96c-43f9-ae6e-60518c9f3238"

type TimetableClient struct {
	APIKey string
}

func NewTimetableClient(key string) *TimetableClient {
	return &TimetableClient{APIKey: key}
}

func (c *TimetableClient) Get() ([]waw.WawValue, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.timetableURL().String())
	if err != nil {
		return nil, err
	}

	var response waw.WawResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	return response.Result, nil
}

func (c *TimetableClient) newRequest() *TimetableRequest {
	return &TimetableRequest{
		ApiKey:     c.APIKey,
		ResourceId: ResourceStopTimetableId,
	}
}

func (c *TimetableClient) timetableURL() *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/dbtimetable_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest().ToValues().Encode()
	return url
}
