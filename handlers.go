package main

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func calculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dataRequest, err := getDataRequestDecode(r.Body)
	if err != nil {
		http.Error(w, getErrorJson(HandlerCalculateHandlerJsonDecodeError.Message), HandlerCalculateHandlerJsonDecodeError.HTTPCode)
		return
	}

	data := dataRequestToData(dataRequest)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var factorialA, factorialB int
	go calculateFactorialWorker(data.a, &wg, &factorialA)
	go calculateFactorialWorker(data.b, &wg, &factorialB)

	wg.Wait()

	dataResponse := getDataResponse(factorialA, factorialB)
	err = getDataRequestEncode(w, dataResponse)
	if err != nil {
		http.Error(w, getErrorJson(HandlerCalculateHandlerJsonEncodeError.Message), HandlerCalculateHandlerJsonEncodeError.HTTPCode)
		return
	}
}

func calculateFactorialWorker(number int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	calculatedFactorial := calculateFactorial(number)
	*result = calculatedFactorial
}
