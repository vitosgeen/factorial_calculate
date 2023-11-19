package usecase

import "factorial_calculate/internal/domain/model"

type IFactorialCalculateUsecase interface {
	Calculate(a, b int) (int, int)
}

type factorialCalculateUsecase struct{}

func NewFactorialCalculateUsecase() IFactorialCalculateUsecase {
	return &factorialCalculateUsecase{}
}

func (f *factorialCalculateUsecase) Calculate(a, b int) (int, int) {
	return model.CalculateFactorial(a), model.CalculateFactorial(b)
}
