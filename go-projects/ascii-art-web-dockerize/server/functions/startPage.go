package functions

import (
	"html/template"
	"log"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		error404(w)
		return
	} else {
		pageTemplate, err := template.ParseFiles("server/templates/index.html")
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		pageTemplate.Execute(w, nil)
	}
}
