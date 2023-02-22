package pages

import (
	"forum/server/data"
	"html/template"
	"net/http"
)

var info []data.RegisterStruct

func Register(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("server/templates/register.html", "server/templates/nav.html")
	if err != nil {
		Error(w, 500)
		return
	}
	if r.Method != http.MethodPost {
		i := data.RegisterStruct{}
		tmpl.Execute(w, i)
		return
	}

	tmpl.Execute(w, info)

}
