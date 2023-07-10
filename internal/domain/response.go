package domain

type Response struct {
	Message       interface{} `json:"message"`
	ErrorResponse string      `json:"error_response"`
	ErrorCode     int         `json:"error_code"`
}
