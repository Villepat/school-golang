package handlers

import (
	"encoding/json"
	"forum/server/database"
	"forum/server/pages"
	"forum/server/session"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pages.Login(w)
		return
	}
	if r.Method == http.MethodPost && database.ValidateUser(r, database.Db) {
		session.StartSesh(w, r)
		json.NewEncoder(w).Encode(&Response{"Login successful", 200})
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Response{"Incorrect username or password", 401})
		return
	}
}
