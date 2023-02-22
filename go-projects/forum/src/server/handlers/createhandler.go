package handlers

import (
	"encoding/json"
	"forum/server/database"
	"forum/server/pages"
	"forum/server/session"
	"net/http"
	"strings"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	//data := pages.MainStruct{}
	if r.Method == "GET" {
		pages.Create(w, r)
		return
	}
	if r.Method == "POST" {
		if logged_in, _ := session.GetSession(w, r); !logged_in {
			pages.Error(w, 401)
			return
		}
		if len(strings.Fields(r.FormValue("title"))) == 0 || len(strings.Fields(r.FormValue("content"))) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{"You cannot\n have empty title or empty content which contains only spaces or tabs", 400})
			return
		}
		title := r.FormValue("title")
		if len(title) < 1 || len(title) > 30 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{"Please create a title consisting of 1-30 characters ", 400})
			return
		}
		content := r.FormValue("content")
		if len(content) < 1 || len(content) > 400 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{"Please provide some content", 400})
			return
		}

		//allow multiple tags: boring, more boring, most boring
		tags := r.FormValue("tags")
		if len(tags) > 30 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{"Max length of tags is 30 characters", 400})
			return
		}

		username, err := session.GetUserName(w, r)
		if err != nil {
			pages.Error(w, 500)
			return
		}
		err = database.AddPost(username, title, content, tags, database.Db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Response{err.Error(), 500})
			return
		}
		json.NewEncoder(w).Encode(&Response{"Post added", 200})
		return
	}
}
