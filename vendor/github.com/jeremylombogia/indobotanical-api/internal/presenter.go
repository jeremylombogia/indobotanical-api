package internal

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
