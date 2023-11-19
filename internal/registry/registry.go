package registry

import (
	"factorial_calculate/internal/interface/controller"
)

type registry struct{}

type Registry interface {
	NewAppController() controller.FactorialCalculateController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.FactorialCalculateController {
	return controller.FactorialCalculateController{
		FactorialCalculateController: r.NewFactorialController(),
	}
}
