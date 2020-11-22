package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "github.com/kdmatrosov/go-rest/models"
	"net/http"
)

// UserHandler is called with localhost:8001/user/:id
type UserHandler struct{}
var userWithoutDb = []m.User{
	{
		Id:    "1",
		Name:  "Kirill",
		Email: "some@email.com",
	},
	{
		Id:    "2",
		Name:  "NotKirill",
		Email: "notsome@email.com",
	},
}
func (handler UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	for _, value := range userWithoutDb {
		if value.Id == id {
			writer.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(writer).Encode(value)
			return
		}
	}

	writer.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"message": "User was not found",
	})
}
