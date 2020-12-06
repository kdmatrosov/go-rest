package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/kdmatrosov/go-rest/db"
	h "github.com/kdmatrosov/go-rest/handlers"
	"log"
	"net/http"
)
func main() {
	db.InitConnection()
	defer db.Conn.Close(context.Background())
	startServer()
}

func startServer() {
	router := mux.NewRouter()
	router.Handle(
		"/",
		h.IndexHandler{},
	)
	router.Handle(
		"/user/{id}",
		h.UserHandler{},
	)
	router.Handle(
		"/user",
		h.AddUserHandler{},
	).Methods("POST")
	port := "8001"
	log.Println("Start server on " + port)
	_ = http.ListenAndServe(":" + port, router)
}
