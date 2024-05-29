package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teixeiragthiago/api-user/config"
	"github.com/teixeiragthiago/api-user/internal/di"
)

func main() {

	config.LoadConfig()

	router := di.SetupDependencies()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), router))
}
