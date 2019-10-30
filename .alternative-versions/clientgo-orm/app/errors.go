package app

import (
	"net/http"

	u "github.com/hfantin/clientgov2/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resources was not found on our server"))
		next.ServeHTTP(w, r)
	})
}
