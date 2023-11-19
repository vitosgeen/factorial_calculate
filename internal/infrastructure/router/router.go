package router

import (
	"factorial_calculate/internal/interface/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(router *httprouter.Router, c controller.FactorialCalculateController) *httprouter.Router {
	router.POST("/calculate", c.FactorialCalculateController.Middleware(c.FactorialCalculateController.CalculateHandler))

	return router
}
