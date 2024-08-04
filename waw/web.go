package waw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WawResponse struct {
	Result []WawValue `json:"result"`
}

type WawValue struct {
	Values []WawArg `json:"values"`
}

type WawArg struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

func (wv WawValue) ToMap() map[string]string {
	result := make(map[string]string)
	for _, value := range wv.Values {
		result[value.Key] = value.Value
	}

	return result
}

func PerformAPIRequest(method string, url string) ([]byte, error) {
	fmt.Println(url)
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte{}))
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
	return resBody, nil
}

func UnmarshalAPIError(body []byte) error {
	var errorResponse ErrorResponse

	err := json.Unmarshal(body, &errorResponse)
	if err != nil {
		return errors.New("unidentified API error occured")
	} else {
		return &WarsawAPIError{ErrorMessage: string(body)}
	}
}
