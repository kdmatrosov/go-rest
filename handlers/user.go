package handlers

import (
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/kdmatrosov/go-rest/db"
	m "github.com/kdmatrosov/go-rest/models"
	"net/http"
)

// UserHandler is called with localhost:8001/user/:id
type UserHandler struct{}

func (handler UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	// скорее всего тут должен быть другой контекст
	row := db.Conn.QueryRow(context.Background(), "select * from users where id=$1", id)
	var user m.User
	err := row.Scan(&user.Id, &user.Email, &user.Name)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(map[string]string{
			"message": "Error during select",
			"err": err.Error(),
		})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(user)
}


type AddUserHandler struct {}
func (handler AddUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var user m.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	sql := "insert into users (email, name) values($1, $2) RETURNING id;"

	// скорее всего тут должен быть другой контекст
	err := db.Conn.QueryRow(context.Background(), sql, user.Email, user.Name).Scan(&user.Id)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(map[string]string{
			"message": "Error during insert",
			"err": err.Error(),
		})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(user)
}
