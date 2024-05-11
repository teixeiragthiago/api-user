package routes

import (
	"github.com/gorilla/mux"
	pingcontroller "github.com/teixeiragthiago/api-user/internal/controller/ping"
)

func RegisterPing(router *mux.Router, pingController *pingcontroller.PingController) {
	router.HandleFunc("/ping", pingController.Ping).Methods("GET")
}
