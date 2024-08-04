package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

const ResourceId = "20d30ed7-fa28-478b-b2e3-17e4a5315e71"

type DictionaryClient struct {
	APIKey string
}

func NewDictionaryClient(key string) *DictionaryClient {
	return &DictionaryClient{APIKey: key}
}

func (c *DictionaryClient) Get() (*Dictionary, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.dictionaryURL().String())
	if err != nil {
		return nil, err
	}

	var response DictionaryResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	return response.Result, nil
}

func (c *DictionaryClient) newRequest() *DictionaryRequest {
	return &DictionaryRequest{
		ApiKey:     c.APIKey,
		ResourceId: ResourceId,
	}
}

func (c *DictionaryClient) dictionaryURL() *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/public_transport_dictionary", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest().ToValues().Encode()
	return url
}
