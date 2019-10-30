package controllers

import (
	"net/http"

	"github.com/hfantin/clientgov2/models"
	u "github.com/hfantin/clientgov2/utils"
)

var GetClients = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetClients()
	// resp := u.Message(true, "success")
	resp := u.SimpleMessage()
	resp["data"] = data
	u.Respond(w, resp)
}
