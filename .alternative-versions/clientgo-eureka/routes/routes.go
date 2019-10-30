package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/hfantin/clientgo/controllers"
)

func Initialize(db *sql.DB) *mux.Router {
	controller := controllers.Controller{}
	controller.Initialize(db)
	router := mux.NewRouter()
	router.HandleFunc("/v1/clients", controller.GetAllClients).Methods("GET")
	router.HandleFunc("/v1/clients/page", controller.GetClients).Methods("GET")
	router.HandleFunc("/v1/clients/{id:[0-9]+}", controller.GetClient).Methods("GET")
	router.HandleFunc("/actuator/info", controllers.Info).Methods("GET")
	router.HandleFunc("/actuator/health", controllers.Health).Methods("GET")
	return router
}
