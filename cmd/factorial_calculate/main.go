package main

import (
	"log"
	"net/http"

	"factorial_calculate/internal/apperrors"
	"factorial_calculate/internal/infrastructure/router"
	"factorial_calculate/internal/registry"

	"github.com/julienschmidt/httprouter"
)

const httpPort = ":8989"

func main() {
	r := httprouter.New()
	reg := registry.NewRegistry()
	routerLocal := router.NewRouter(r, reg.NewAppController())
	err := http.ListenAndServe(httpPort, routerLocal)
	if err != nil {
		log.Fatal(apperrors.MainListenAndServe.Message)
	}
}
