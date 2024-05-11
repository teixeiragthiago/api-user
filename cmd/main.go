package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teixeiragthiago/api-user/internal/di"
)

func main() {

	router := di.SetupControllers()

	fmt.Print("Running API on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
