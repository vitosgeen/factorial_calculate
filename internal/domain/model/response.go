package model

type DataResponse struct {
	A int `json:"a"`
	B int `json:"b"`
}

func MapDataToDataResponse(data *Data) *DataResponse {
	return &DataResponse{
		A: data.A,
		B: data.B,
	}
}
