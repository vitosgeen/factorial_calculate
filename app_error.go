package main

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

func getErrorJson(strError string) string {
	errString, _ := json.Marshal(getErrorResponse(strError))
	return string(errString)
}

func getErrorResponse(strError string) *errorResponse {
	return &errorResponse{
		Error: strError,
	}
}
