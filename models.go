package main

type data struct {
	a int
	b int
}

type dataRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type dataResponse struct {
	A int `json:"a"`
	B int `json:"b"`
}

func dataRequestToData(dataRequest *dataRequest) *data {
	return &data{
		a: dataRequest.A,
		b: dataRequest.B,
	}
}

func getDataResponse(a int, b int) *dataResponse {
	return &dataResponse{
		A: a,
		B: b,
	}
}

func calculateFactorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= number; i++ {
		result *= i
	}
	return result
}
