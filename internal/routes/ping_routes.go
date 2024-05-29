package routes

import (
	"github.com/gin-gonic/gin"
	pingcontroller "github.com/teixeiragthiago/api-user/internal/controller/ping"
)

// func RegisterPing(router *mux.Router, pingController *pingcontroller.PingController) {
// 	router.HandleFunc("/ping", pingController.Ping).Methods("GET")
// }

func RegisterPing(r *gin.Engine, pingController *pingcontroller.PingController) {
	r.GET("/ping", pingController.Ping)
}
