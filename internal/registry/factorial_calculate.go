package registry

import (
	"factorial_calculate/internal/interface/controller"
	"factorial_calculate/internal/usecase/usecase"
)

func (r *registry) NewFactorialController() controller.IFactorialCalculateController {
	factorialUsecase := usecase.NewFactorialCalculateUsecase()

	return controller.NewFactorialCalculateController(factorialUsecase)
}
