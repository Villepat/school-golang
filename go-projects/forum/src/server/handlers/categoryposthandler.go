package handlers

import (
	"forum/server/pages"
	"net/http"
)

func CategoryPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pages.CategoryPosts(r, w)
		return
	}
}
