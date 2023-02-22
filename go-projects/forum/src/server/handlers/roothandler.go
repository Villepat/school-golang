package handlers

import (
	"forum/server/pages"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pages.Home(r, w)
		return
	}
}
