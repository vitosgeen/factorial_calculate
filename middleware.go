package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	dataRequest := &dataRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, dataRequest)
	if err != nil {
		return err
	}

	if dataRequest.A < 0 || dataRequest.B < 0 {
		return fmt.Errorf("a and b must be positive")
	}

	r.Body = io.NopCloser(bytes.NewReader(body))

	return nil
}
