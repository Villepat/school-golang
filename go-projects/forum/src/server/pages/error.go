package pages

import (
	"fmt"
	"forum/server/data"
	"net/http"
	"text/template"
)

func Error(w http.ResponseWriter, error int) {
	msg := data.ErrorStruct{}
	msg.Data = "Non-recognizable error"
	if error == 400 {
		msg.Data = "400 Bad Request"
	}
	if error == 404 {
		msg.Data = "404 not found"
	}
	if error == 401 {
		msg.Data = "Unauthorized user."
	}
	if error == 500 {
		msg.Data = "500 internal server error"
	}
	if error == 405 {
		msg.Data = "405 method not allowed"
	}
	if error == 666 {
		msg.Data = "User already exists!"
	}
	tmpl, err := template.ParseFiles("server/templates/error.html")
	if err != nil {
		fmt.Fprintf(w, "500 internal server error")
	}
	tmpl.Execute(w, msg)
}
