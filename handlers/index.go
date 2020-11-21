package handlers

import "net/http"

// IndexHandler is called with localhost:8001/
type IndexHandler struct{}

func (handler IndexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("index"))
}
