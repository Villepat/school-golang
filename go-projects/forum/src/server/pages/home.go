package pages

import (
	"fmt"
	datapackage "forum/server/data"
	"forum/server/database"
	"forum/server/session"
	"net/http"
	"strings"
	"text/template"
)

func Home(r *http.Request, w http.ResponseWriter) {

	data := datapackage.MainStruct{}
	posts, err := database.GetPosts("")
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}

	for i := range posts {
		posts[i], _ = database.GetAmountOfLikes(posts[i])
	}
	data.Data = datapackage.HomeStruct{Topics: posts}
	data.LoggedIn, data.Username = session.GetSession(w, r)
	// remove comma and unneccecary spaces from tags
	for i := 0; i < len(data.Data.(datapackage.HomeStruct).Topics); i++ {
		data.Data.(datapackage.HomeStruct).Topics[i].Tags = strings.Replace(data.Data.(datapackage.HomeStruct).Topics[i].Tags, ",", " ", -1)
		data.Data.(datapackage.HomeStruct).Topics[i].Tags = strings.Replace(data.Data.(datapackage.HomeStruct).Topics[i].Tags, "  ", " ", -1)

	}
	tmpl, err := template.ParseFiles("server/templates/home.html", "server/templates/nav.html")
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}
	tmpl.Execute(w, data)
}
