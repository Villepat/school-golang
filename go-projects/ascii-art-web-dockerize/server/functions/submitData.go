package functions

import (
	"asciiartweb/server/ascii"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var DataToPage struct {
	AsciiArt    string
	PlaceHolder string
}

func SubmitData(w http.ResponseWriter, r *http.Request) {
	banner := ""
	input := ""
	if r.Method != "POST" && r.URL.Path != "/ascii-art" {
		error400(w)
	} else if r.Method == "POST" && r.URL.Path == "/ascii-art" {
		tmpl, err := template.ParseFiles("server/templates/index.html")

		r.ParseForm()
		if r.FormValue("banner_style") != "standard" &&
			r.FormValue("banner_style") != "thinkertoy" &&
			r.FormValue("banner_style") != "shadow" {
			error400(w)
			return
		}

		banner = r.FormValue("banner_style")

		input = r.FormValue("inputString")
		inputInter := strings.Split(input, "\r\n")
		resultString := ""
		for i, val := range inputInter {
			if i < len(inputInter)-1 {
				resultString += val + "\\n"
			} else {
				resultString += val
			}
		}

		DataToPage.AsciiArt = ascii.AsciiPrinter(resultString, banner)
		DataToPage.PlaceHolder = input

		w.WriteHeader(200)
		tmpl.Execute(w, DataToPage)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		error500(w)
	}
}
