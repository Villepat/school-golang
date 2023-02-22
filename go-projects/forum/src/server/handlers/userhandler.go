package handlers

import (
	"forum/server/pages"
	"forum/server/session"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if logged_in, _ := session.GetSession(w, r); !logged_in {
		pages.Error(w, 401)
		return
	}
	if r.Method == "GET" {
		pages.UserPage(r, w)
		return
	}
}
