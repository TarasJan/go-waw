package waw

import "fmt"

type ErrorResponse struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type WarsawAPIError struct {
	ErrorMessage string
}

func (e *WarsawAPIError) Error() string {
	return fmt.Sprintf("Warsaw API failed to process request. Maek sure that you have correct API key set in GOWAW_KEY env var. Response: %s", e.ErrorMessage)
}
