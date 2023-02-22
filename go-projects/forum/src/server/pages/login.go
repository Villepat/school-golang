package pages

import (
	"forum/server/data"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter) {
	info := data.LoginStruct{Data: "Login"}
	tmpl, err := template.ParseFiles("server/templates/login.html", "server/templates/nav.html")
	if err != nil {
		Error(w, 500)
		return
	}
	tmpl.Execute(w, info)
}
