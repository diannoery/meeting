package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
func RunServer(router *mux.Router) {
	port := "localhost:8080"
	fmt.Println("Setting Web Server at port : " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
