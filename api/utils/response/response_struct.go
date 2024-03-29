package response

import (
	"api-rest/api/utils/types"
)

func SuccessResponse(message string, data interface{}) types.SuccessResponse {
	if message == "" {
		message = "Ok"

	}
	return types.SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, error interface{}) types.ErrorResponse {
	if message == "" {
		message = "Error"
	}
	return types.ErrorResponse{
		Status:  false,
		Message: message,
		Erorr:   error,
	}
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
