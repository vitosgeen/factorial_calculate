package controller

import (
	"fmt"
	"net/http"

	"factorial_calculate/internal/apperrors"
	"factorial_calculate/internal/domain/model"
	"factorial_calculate/internal/utils"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func (f *factorialCalculateController) Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := validateRequest(w, r, ps)
		if err != nil {
			http.Error(w, apperrors.GetErrorJson(apperrors.MiddlewareIncorrectInputError.Message), apperrors.MiddlewareIncorrectInputError.HTTPCode)
			return
		}
		next(w, r, ps)
	}
}

func validateRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	body, err := utils.ReadRequestBody(r)
	if err != nil {
		return err
	}

	dataRequest, err := utils.UnmarshalRequestBody(body)
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

	utils.ResetRequestBody(r, body)
	return nil
}

func validateDataRequestStruct(dataReq model.DataRequest) error {
	validate := validator.New()
	err := validate.Struct(dataReq)
	if err != nil {
		return err
	}
	return nil
}

func validateDataRequest(dataReq model.DataRequest) error {
	if dataReq.A < 0 || dataReq.B < 0 {
		return fmt.Errorf("a and b must be positive")
	}
	return nil
}
