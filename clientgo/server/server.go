package server

import (
	"log"
	"net/http"

	"github.com/hfantin/clientgo/db"
	r "github.com/hfantin/clientgo/routes"
	"github.com/hfantin/clientgo/utils"
)

func StartWebServer(env utils.Environment) {
	db := db.Connect(env.DBHost, env.DBUser, env.DBPass, env.DBName)
	// initialize router
	router := r.Initialize(db)
	log.Println("Serving on port", env.ServerPort)
	err := http.ListenAndServe(":"+env.ServerPort, router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port", env.ServerPort)
		log.Println("Error: " + err.Error())
	}
}
