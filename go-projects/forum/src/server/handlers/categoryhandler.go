package handlers

import (
	"forum/server/pages"
	"net/http"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pages.Categories(r, w)
		return
	}
}
