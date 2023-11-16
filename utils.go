package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func getDataRequestEncode(w http.ResponseWriter, dataResponse *dataResponse) error {
	err := json.NewEncoder(w).Encode(dataResponse)
	if err != nil {
		return err
	}

	return nil
}

func getDataRequestDecode(body io.ReadCloser) (*dataRequest, error) {
	dataReq := &dataRequest{}
	err := json.NewDecoder(body).Decode(dataReq)
	if err != nil {
		return nil, err
	}

	return dataReq, nil
}

func readRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func unmarshalRequestBody(body []byte) (dataRequest, error) {
	var dataBody dataRequest
	err := json.Unmarshal(body, &dataBody)
	if err != nil {
		return dataRequest{}, err
	}
	return dataBody, nil
}

func resetRequestBody(r *http.Request, body []byte) {
	r.Body = io.NopCloser(bytes.NewBuffer(body))
}
