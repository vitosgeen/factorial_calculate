package model

type DataRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

func MapDataRequestToData(dataRequest *DataRequest) *Data {
	return &Data{
		A: dataRequest.A,
		B: dataRequest.B,
	}
}
