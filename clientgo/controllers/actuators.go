package controllers

import (
	"encoding/json"

	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	status := make(map[string]interface{})
	status["status"] = "OK"
	if err := json.NewEncoder(w).Encode(status); err != nil {
		panic(err)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	health := make(map[string]interface{})
	health["health"] = "OK"
	if err := json.NewEncoder(w).Encode(health); err != nil {
		panic(err)
	}
}
