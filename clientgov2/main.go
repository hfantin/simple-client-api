package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hfantin/clientgov2/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/clients", controllers.GetClients).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
