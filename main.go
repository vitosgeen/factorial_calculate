package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const httpPort = ":8989"

func main() {
	router := httprouter.New()
	router.POST("/calculate", middleware(calculateHandler))
	log.Fatal(http.ListenAndServe(httpPort, router))
}
