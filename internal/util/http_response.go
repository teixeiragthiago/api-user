package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpResponseErrorHandler interface {
	Error(w http.ResponseWriter, statusCode int, err error)
	Success(w http.ResponseWriter, statusCode int, data interface{})
}

type httpResponse struct {
}

func NewResponseErrorHandler() HttpResponseErrorHandler {
	return &httpResponse{}
}

func (e *httpResponse) Error(w http.ResponseWriter, statusCode int, err error) {

	jsonContent(w, statusCode, struct {
		ErrorMessage string `json:"ErrorMessage"`
	}{
		ErrorMessage: err.Error(),
	})
}

func (e *httpResponse) Success(w http.ResponseWriter, statusCode int, data interface{}) {

	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

func jsonContent(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
