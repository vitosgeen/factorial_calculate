package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", middleware(calculateHandler))
	log.Fatal(http.ListenAndServe(":8989", router))
}
