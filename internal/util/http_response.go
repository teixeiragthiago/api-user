package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpResponseErrorHandler interface {
	Error(w http.ResponseWriter, statusCode int, err error)
	Success(w http.ResponseWriter, statusCode int, err error)
}

type httpResponse struct {
}

func (e *httpResponse) Success(w http.ResponseWriter, statusCode int, err error) {
	panic("unimplemented")
}

func NewResponseErrorHandler() HttpResponseErrorHandler {
	return &httpResponse{}
}

func (e *httpResponse) Error(w http.ResponseWriter, statusCode int, err error) {

	jsonError(w, statusCode, struct {
		ErrorMessage string `json:"ErrorMessage"`
	}{
		ErrorMessage: err.Error(),
	})
}

func jsonError(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
