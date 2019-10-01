package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hfantin/clientgo/model"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/v1/clients", app.getAllClients).Methods("GET")
	app.Router.HandleFunc("/v1/clients/page", app.getClients).Methods("GET")
	app.Router.HandleFunc("/v1/clients/{id:[0-9]+}", app.getClient).Methods("GET")
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Println("    ______               _      _____  _____ ")
	fmt.Println("   / ____/ (_)__  ____  / /_   / ____/ __  /")
	fmt.Println("  / /   / / / _  / __  / __/  / / __/ / / /")
	fmt.Println(" / /___/ / /  __/ / / / /_   / /_/ / /_/ / ")
	fmt.Println("/_____/_/_/____/_/ /_/ __/   _____/ ____/ ")
	fmt.Println("")
	log.Println("connecting to database...")
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.DB.SetMaxOpenConns(10)
	a.DB.SetMaxIdleConns(5)
	// a.DB.SetConnMaxLifetime(1000)
	log.Println("creating routes...")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Println("Starting server on port " + addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}
	client := model.Client{ID: id}
	if err := client.GetClient(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Client not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, client)
}

func (app *App) getClients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	clients, err := model.GetClients(app.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, clients)
}

func (app *App) getAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := model.GetAllClients(app.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, clients)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
