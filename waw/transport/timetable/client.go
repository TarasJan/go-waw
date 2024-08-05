package timetable

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
	"github.com/TarasJan/go-waw/waw/transport/stop"
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

func (c *TimetableClient) GetTimetableFor(stop *stop.Stop, line string) (*Timetable, error) {
	request := c.newRequest(ResourceTimetableId).WithStop(stop).WithLine(line)
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.timetableURL(request).String())
	if err != nil {
		return nil, err
	}

	var response waw.WawResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	timetable, err := NewTimetableFrom(response)
	if err != nil {
		return nil, err
	}

	return timetable.WithLine(line), nil
}

func (c *TimetableClient) GetLinesFor(stop *stop.Stop) ([]string, error) {
	request := c.newRequest(ResourceLineTimetableId).WithStop(stop)
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.timetableURL(request).String())
	if err != nil {
		return nil, err
	}

	var response waw.WawResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	lines := make([]string, 0)
	for _, record := range response.Result {
		lines = append(lines, record.Values[0].Value)
	}

	return lines, nil

}

func (c *TimetableClient) Get() ([]waw.WawValue, error) {
	request := c.newRequest(ResourceLineTimetableId)
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.timetableURL(request).String())
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

func (c *TimetableClient) newRequest(resourceId string) *TimetableRequest {
	return &TimetableRequest{
		ApiKey:     c.APIKey,
		ResourceId: resourceId,
	}
}

func (c *TimetableClient) timetableURL(request *TimetableRequest) *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/dbtimetable_get", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = request.ToValues().Encode()
	return url
}
