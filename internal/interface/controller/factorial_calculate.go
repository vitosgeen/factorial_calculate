package controller

import (
	"net/http"
	"sync"

	"factorial_calculate/internal/apperrors"
	"factorial_calculate/internal/domain/model"
	"factorial_calculate/internal/usecase/usecase"
	"factorial_calculate/internal/utils"

	"github.com/julienschmidt/httprouter"
)

type IFactorialCalculateController interface {
	CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Middleware(next httprouter.Handle) httprouter.Handle
}

type FactorialCalculateController struct {
	FactorialCalculateController IFactorialCalculateController
}

type factorialCalculateController struct {
	factoctorialUsecase usecase.IFactorialCalculateUsecase
}

func NewFactorialCalculateController(factorialUsecase usecase.IFactorialCalculateUsecase) IFactorialCalculateController {
	return &factorialCalculateController{
		factoctorialUsecase: factorialUsecase,
	}
}

func (f *factorialCalculateController) CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dataRequest, err := utils.GetDataRequestDecode(r.Body)
	if err != nil {
		http.Error(w, apperrors.GetErrorJson(apperrors.HandlerCalculateHandlerJsonDecodeError.Message), apperrors.HandlerCalculateHandlerJsonDecodeError.HTTPCode)
		return
	}

	data := model.MapDataRequestToData(dataRequest)
	runWaitGroupFactorialCalculate(data)

	dataResponse := model.MapDataToDataResponse(data)

	err = utils.GetDataRequestEncode(w, dataResponse)
	if err != nil {
		http.Error(w, apperrors.GetErrorJson(apperrors.HandlerCalculateHandlerJsonEncodeError.Message), apperrors.HandlerCalculateHandlerJsonEncodeError.HTTPCode)
		return
	}
}

func runWaitGroupFactorialCalculate(data *model.Data) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var factorialA, factorialB int
	go calculateFactorialWorker(data.A, &wg, &factorialA)
	go calculateFactorialWorker(data.B, &wg, &factorialB)
	wg.Wait()

	data.A = factorialA
	data.B = factorialB
}

func calculateFactorialWorker(number int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	calculatedFactorial := model.CalculateFactorial(number)
	*result = calculatedFactorial
}
