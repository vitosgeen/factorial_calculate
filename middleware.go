package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := validateRequest(w, r, ps)
		if err != nil {
			http.Error(w, getErrorJson(MiddlewareIncorrectInputError.Message), MiddlewareIncorrectInputError.HTTPCode)
			return
		}
		next(w, r, ps)
	}
}

func validateRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	body, err := readRequestBody(r)
	if err != nil {
		return err
	}

	dataRequest, err := unmarshalRequestBody(body)
	if err != nil {
		return err
	}

	err = validateDataRequest(dataRequest)
	if err != nil {
		return err
	}

	err = validateDataRequestStruct(dataRequest)
	if err != nil {
		return err
	}

	resetRequestBody(r, body)
	return nil
}

func validateDataRequestStruct(dataReq dataRequest) error {
	validate := validator.New()
	err := validate.Struct(dataReq)
	if err != nil {
		return err
	}
	return nil
}

func validateDataRequest(dataReq dataRequest) error {
	if dataReq.A < 0 || dataReq.B < 0 {
		return fmt.Errorf("a and b must be positive")
	}
	return nil
}
