package handlers

import (
	"forum/server/session"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session.Logout(w, r)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
