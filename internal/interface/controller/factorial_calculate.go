package controller

import (
	"net/http"

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
	data = f.factoctorialUsecase.RunWaitGroupFactorialCalculate(data)

	dataResponse := model.MapDataToDataResponse(data)

	err = utils.GetDataRequestEncode(w, dataResponse)
	if err != nil {
		http.Error(w, apperrors.GetErrorJson(apperrors.HandlerCalculateHandlerJsonEncodeError.Message), apperrors.HandlerCalculateHandlerJsonEncodeError.HTTPCode)
		return
	}
}
