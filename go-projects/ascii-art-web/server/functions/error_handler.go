package functions

import (
	"net/http"
)

// Handles 400 bad request
func error400(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte("400 Bad request"))
}

// Handles 404 not found
func error404(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("404 Page not found"))
}

// Handles 500 internal server error
func error500(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("500 Internal Server Error"))
}
