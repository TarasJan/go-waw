package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/TarasJan/go-waw/waw"
)

const ResourceId = "b61f1869-6f0a-4375-acdf-87ccffeecdf0"

type RouteClient struct {
	APIKey string
}

func NewRouteClient(key string) *RouteClient {
	return &RouteClient{APIKey: key}
}

func (c *RouteClient) Get() (LineRoutes, error) {
	resBody, err := waw.PerformAPIRequest(http.MethodGet, c.routesURL().String())
	if err != nil {
		return nil, err
	}

	var response RoutesResponse

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, waw.UnmarshalAPIError(resBody)
	}

	return response.Result, nil
}

func (c *RouteClient) newRequest() *RoutesRequest {
	return &RoutesRequest{
		ApiKey: c.APIKey,
	}
}

func (c *RouteClient) routesURL() *url.URL {
	urlBase := fmt.Sprintf("%s/api/action/public_transport_routes", waw.APIURL)
	url, _ := url.Parse(urlBase)
	url.RawQuery = c.newRequest().ToValues().Encode()
	return url
}
