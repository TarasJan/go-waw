package dictionary

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

const ResourceId = "20d30ed7-fa28-478b-b2e3-17e4a5315e71"

type DictionaryClient struct {
	APIKey string
}

func NewDictionaryClient(key string) *DictionaryClient {
	return &DictionaryClient{APIKey: key}
}

func (c *DictionaryClient) Get() (*Dictionary, error) {
	req, err := http.NewRequest(http.MethodGet, c.dictionaryURL().String(), bytes.NewReader([]byte{}))
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

	var response DictionaryResponse
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
