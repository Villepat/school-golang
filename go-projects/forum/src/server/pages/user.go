package pages

import (
	"fmt"
	datapackage "forum/server/data"
	"forum/server/database"
	"forum/server/session"
	"net/http"
	"text/template"
)

func UserPage(r *http.Request, w http.ResponseWriter) {
	data := datapackage.MainStruct{}
	data.LoggedIn, data.Username = session.GetSession(w, r)
	posts, err := database.GetPostsUser(data.Username)
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}

	likeposts, err := database.GetPostsLikedByUser(data.Username)
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}
	commentposts, err := database.GetPostThroughLikedComments(data.Username)
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}

	data.Data = datapackage.UserProfile{Data: database.GetUserProfile(data.Username), Posts: posts, Likes: likeposts, CommentLikes: commentposts}
	tmpl, err := template.ParseFiles("server/templates/user.html", "server/templates/nav.html")
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}
	tmpl.Execute(w, data)

}
