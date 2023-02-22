package handlers

import (
	"forum/server/database"
	"forum/server/pages"
	"forum/server/session"
	"net/http"
	"strconv"
	"strings"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	for _, v := range pages.AllowedURLs() {
		if v == r.URL.Path && r.Method == "GET" {
			pages.Post(w, r)
			return
		}
		if v == r.URL.Path && r.Method == "POST" {
			if logged_in, _ := session.GetSession(w, r); !logged_in {
				pages.Error(w, 401)
				return
			}
			username, err := session.GetUserName(w, r)
			if err != nil {
				pages.Error(w, 500)
				return
			}
			content := r.FormValue("content")
			if content == "" || len(strings.Fields(r.FormValue("content"))) == 0 {
				http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
				return
			}
			post_id, err := strconv.Atoi(strings.Split(v[1:], "/")[1])
			if err != nil {
				pages.Error(w, 500)
				return
			}
			database.AddComment(username, content, post_id)
			return
		}
		if v == r.URL.Path && r.Method == "DELETE" {
			post_id, err := strconv.Atoi(strings.Split(v[1:], "/")[1])
			if err != nil {
				pages.Error(w, 500)
				return
			}
			database.DeletePost(post_id, database.Db)
			return
		}
	}
}
