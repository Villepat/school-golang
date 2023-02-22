package pages

import (
	"forum/server/data"
	"forum/server/database"
	"forum/server/session"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Post(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetPosts("")
	if err != nil {
		log.Println(err)
		Error(w, 500)
		return
	}

	data := data.MainStruct{}
	var post_id int

	data.LoggedIn, data.Username = session.GetSession(w, r)
	for _, v := range AllowedURLs() {
		if v == r.URL.Path {
			words := strings.Split(v[1:], "/")
			post_id, err = strconv.Atoi(words[1])
			if err != nil {
				Error(w, 404)
				return
			}
			database.IncrementViews(post_id, database.Db)
			current_post := database.GetSpecificPost(posts, post_id)
			current_post.Comments = database.GetComments(post_id)
			current_post, err = database.GetLikes(data.Username, database.Db, current_post)
			if err != nil {
				Error(w, 500)
				return
			}
			current_post, err = database.GetAmountOfLikes(current_post)
			if err != nil {
				Error(w, 500)
				return
			}
			// set Username and LoggedIn for each comment
			for i := range current_post.Comments {
				current_post.Comments[i].Username = data.Username
				current_post.Comments[i].LoggedIn = data.LoggedIn
			}
			data.Data = current_post
			break
		}
	}

	data.Tags, err = database.GetTags(post_id)
	if err != nil {
		Error(w, 500)
	}
	tmpl, err := template.ParseFiles("server/templates/post.html", "server/templates/nav.html")
	if err != nil {
		Error(w, 500)
	}
	tmpl.Execute(w, data)
}
