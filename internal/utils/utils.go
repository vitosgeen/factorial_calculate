package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"factorial_calculate/internal/domain/model"
)

func GetDataRequestEncode(w http.ResponseWriter, dataResponse *model.DataResponse) error {
	err := json.NewEncoder(w).Encode(dataResponse)
	if err != nil {
		return err
	}

	return nil
}

func GetDataRequestDecode(body io.ReadCloser) (*model.DataRequest, error) {
	dataReq := &model.DataRequest{}
	err := json.NewDecoder(body).Decode(dataReq)
	if err != nil {
		return nil, err
	}

	return dataReq, nil
}

func ReadRequestBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func UnmarshalRequestBody(body []byte) (model.DataRequest, error) {
	var dataBody model.DataRequest
	err := json.Unmarshal(body, &dataBody)
	if err != nil {
		return model.DataRequest{}, err
	}
	return dataBody, nil
}

func ResetRequestBody(r *http.Request, body []byte) {
	r.Body = io.NopCloser(bytes.NewBuffer(body))
}
