package pingcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

// func (c *PingController) Ping(w http.ResponseWriter, r *http.Request) {

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Pong! API is alive"))
// }

func (p *PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
