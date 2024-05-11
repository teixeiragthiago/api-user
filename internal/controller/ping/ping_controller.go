package pingcontroller

import "net/http"

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (c *PingController) Ping(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong! API is alive"))
}
