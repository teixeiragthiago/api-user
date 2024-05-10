package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// userController, err := di.SetupDependencies()
	// if err != nil {
	// 	log.Fatal("Error setting up dependencies", err)
	// }

	router := mux.NewRouter()

	// routes.SetupRoutes(router, userController)

	fmt.Print("Running API on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
