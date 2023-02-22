package pages

import (
	"forum/server/data"
	"forum/server/session"
	"net/http"
	"text/template"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := data.MainStruct{}
	data.LoggedIn, data.Username = session.GetSession(w, r)
	if !data.LoggedIn {
		Error(w, 401)
		return
	}
	tmpl, err := template.ParseFiles("server/templates/create.html", "server/templates/nav.html")
	if err != nil {
		Error(w, 500)
		return
	}
	tmpl.Execute(w, data)
}
