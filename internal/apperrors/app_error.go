package apperrors

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

type AppError struct {
	Message  string
	Code     string
	HTTPCode int
}

var (
	MainListenAndServe = AppError{
		Message:  "Failed to listen and serve",
		Code:     "LISTEN_AND_SERVE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}

	MiddlewareIncorrectInputError = AppError{
		Message:  "Incorrect input",
		Code:     "INCORRECT_INPUT",
		HTTPCode: http.StatusBadRequest,
	}

	HandlerCalculateHandlerJsonDecodeError = AppError{
		Message:  "Failed to decode json",
		Code:     "JSON_DECODE_ERR",
		HTTPCode: http.StatusBadRequest,
	}

	HandlerCalculateHandlerJsonEncodeError = AppError{
		Message:  "Failed to encode json",
		Code:     "JSON_ENCODE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
)

func GetErrorJson(strError string) string {
	errString, _ := json.Marshal(GetErrorResponse(strError))
	return string(errString)
}

func GetErrorResponse(strError string) *errorResponse {
	return &errorResponse{
		Error: strError,
	}
}
