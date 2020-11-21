package main

import (
	"github.com/gorilla/mux"
	"github.com/kdmatrosov/go-rest/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Handle(
		"/",
		handlers.IndexHandler{},
	)
	port := "8001"
	log.Println("Start server on " + port)
	_ = http.ListenAndServe(":" + port, router)
}
