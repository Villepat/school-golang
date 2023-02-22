package server

import (
	"html/template"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" || r.Method != "GET" {
		error404(w)
		return
	} else {
		tmpl, err := template.ParseFiles("Front-end/index.html")
		if err != nil {
			error500(w)
			return
		}
		err2 := tmpl.ExecuteTemplate(w, "index.html", PageData(w))
		if err2 != nil {
			error500(w)
			return
		}
	}
}
