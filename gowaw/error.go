package gowaw

type ErrorResponse struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type UnauthorizedAccessError struct{}

func (e *UnauthorizedAccessError) Error() string {
	return "Warsaw API failed to recognize the provided API token"
}
