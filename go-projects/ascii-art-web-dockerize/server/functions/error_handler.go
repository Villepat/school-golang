package functions

import (
	"net/http"
)

func error400(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Write([]byte("400 Bad request"))
}

func error404(w http.ResponseWriter) {
	w.WriteHeader(404)
	w.Write([]byte("404 Page not found"))
}

func error500(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("500 Internal Server Error"))
}
