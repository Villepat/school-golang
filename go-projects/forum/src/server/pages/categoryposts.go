package pages

import (
	"fmt"
	datapackage "forum/server/data"
	"forum/server/database"
	"forum/server/session"
	"net/http"
	"text/template"
)

func CategoryPosts(r *http.Request, w http.ResponseWriter) {
	data := datapackage.MainStruct{}
	topics, err := database.GetPosts(r.URL.Path[12:])
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}
	data.Data = datapackage.HomeStruct{Topics: topics}
	data.LoggedIn, data.Username = session.GetSession(w, r)
	tmpl, err := template.ParseFiles("server/templates/home.html", "server/templates/nav.html")
	if err != nil {
		fmt.Println(err)
		Error(w, 500)
		return
	}
	tmpl.Execute(w, data)

}
