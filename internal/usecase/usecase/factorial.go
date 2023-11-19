package usecase

import (
	"sync"

	"factorial_calculate/internal/domain/model"
)

type IFactorialCalculateUsecase interface {
	RunWaitGroupFactorialCalculate(data *model.Data) *model.Data
}

type factorialCalculateUsecase struct{}

func NewFactorialCalculateUsecase() IFactorialCalculateUsecase {
	return &factorialCalculateUsecase{}
}

func (f *factorialCalculateUsecase) RunWaitGroupFactorialCalculate(data *model.Data) *model.Data {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var factorialA, factorialB int
	go calculateFactorialWorker(data.A, &wg, &factorialA)
	go calculateFactorialWorker(data.B, &wg, &factorialB)
	wg.Wait()

	data.A = factorialA
	data.B = factorialB

	return data
}

func calculateFactorialWorker(number int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	calculatedFactorial := model.CalculateFactorial(number)
	*result = calculatedFactorial
}
