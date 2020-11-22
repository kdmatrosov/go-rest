package main

import (
	"github.com/gorilla/mux"
	h "github.com/kdmatrosov/go-rest/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Handle(
		"/",
		h.IndexHandler{},
	)
	router.Handle(
		"/user/{id}",
		h.UserHandler{},
	)
	port := "8001"
	log.Println("Start server on " + port)
	_ = http.ListenAndServe(":" + port, router)
}
