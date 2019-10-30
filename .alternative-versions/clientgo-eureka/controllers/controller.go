package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hfantin/clientgo/model"
	u "github.com/hfantin/clientgo/utils"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) Initialize(db *sql.DB) {
	c.DB = db
}

func (c *Controller) GetAllClients(w http.ResponseWriter, r *http.Request) {
	clients, err := model.GetAllClients(c.DB)
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJSON(w, http.StatusOK, clients)
}

func (c *Controller) GetClients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	clients, err := model.GetClients(c.DB, start, count)
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJSON(w, http.StatusOK, clients)
}

func (c *Controller) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}
	client := model.Client{ID: id}
	if err := client.GetClient(c.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			u.RespondWithError(w, http.StatusNotFound, "Client not found")
		default:
			u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	u.RespondWithJSON(w, http.StatusOK, client)
}
