package handlers

import (
	"fmt"
	"forum/server/database"
	"forum/server/pages"
	"forum/server/session"
	"net/http"
	"strconv"
	"strings"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	array := strings.Split(r.URL.Path, "/")
	status := 0
	username, err := session.GetUserName(w, r)
	if err != nil {
		fmt.Println(err.Error())
		pages.Error(w, 500)
		return
	}
	if username == "" && err == nil {
		pages.Error(w, 401)
		return
	}
	post_id := 0
	comment_id := 0
	if array[1] == "liked" {
		status = 1
	}
	if array[1] == "disliked" {
		status = -1
	}
	if status == 0 {
		pages.Error(w, 400)
		return
	}
	num, err := strconv.Atoi(array[2])
	if err != nil {
		pages.Error(w, 500)
		return
	}
	if array[3] == "post" {
		post_id = num
	}
	if array[3] == "comment" {
		comment_id = num
	}

	if err := database.AddLikeDislike(status, username, post_id, comment_id, database.Db); err != nil {
		fmt.Println(err.Error())
		pages.Error(w, 500)
		return
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
