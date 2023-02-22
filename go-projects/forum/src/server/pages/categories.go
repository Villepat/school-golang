package pages

import (
	"fmt"
	datapackage "forum/server/data"
	"forum/server/database"
	"forum/server/session"
	"net/http"
	"text/template"
)

func Categories(r *http.Request, w http.ResponseWriter) {

	data := datapackage.MainStruct{}

	data.LoggedIn, data.Username = session.GetSession(w, r)
	tmpl, err := template.ParseFiles("server/templates/categories.html", "server/templates/nav.html")
	if err != nil {
		Error(w, 500)
		fmt.Println(err)
		return
	}

	// Convert the tags data to a JavaScript array literal
	tags, err := database.GetTags(0)
	if err != nil {
		Error(w, 500)
		fmt.Println(err)
		return
	}
	tagsJS := "["
	for i, tag := range tags {
		tagsJS += "\"" + tag + "\""
		if i < len(tags)-1 {
			tagsJS += ", "
		}
	}
	tagsJS += "]"
	// add tagsJS to the data struct
	data.Data = datapackage.CategoryStruct{TagsJS: tagsJS, Tags: tags}

	tmpl.Execute(w, data)
}
