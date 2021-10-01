package models

type JSONResponse struct {
	Status       bool        `json:"status"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}

func SuccessResponse(response JSONResponse) JSONResponse {
	response.Status = true
	return response
}

func ErrorResponse(response JSONResponse, message string) JSONResponse {
	response.Status = false
	response.ErrorMessage = message
	return response
}