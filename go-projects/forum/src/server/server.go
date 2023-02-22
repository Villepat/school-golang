package server

import (
	"fmt"
	"forum/server/handlers"
	"forum/server/pages"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", handler)
	css := http.FileServer(http.Dir("server/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	js := http.FileServer(http.Dir("server/js/"))
	http.Handle("/js/", http.StripPrefix("/js/", js))
	imgs := http.FileServer(http.Dir("server/imgs/"))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", imgs))
	tmp := http.FileServer(http.Dir("server/template-helpers"))
	http.Handle("/template-helpers/", http.StripPrefix("/template-helpers/", tmp))
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for _, v := range pages.AllowedURLs() {
		if r.URL.Path == v {
			handlers.PostsHandler(w, r)
			return
		}
	}
	for _, v := range pages.AllowedCategoriesURLs() {
		if r.URL.Path == v {
			handlers.CategoryPostHandler(w, r)
			return
		}
	}
	if r.URL.Path == "/" {
		handlers.RootHandler(w, r)
		return
	}
	if r.URL.Path == "/create" {
		handlers.CreateHandler(w, r)
		return
	}
	if r.URL.Path == "/logout" {
		handlers.LogoutHandler(w, r)
		return
	}
	if r.URL.Path == "/register" {
		handlers.RegisterHandler(w, r)
		return
	}
	if r.URL.Path == "/login" {
		handlers.LoginHandler(w, r)
		return
	}
	if r.URL.Path == "/user" {
		handlers.UserHandler(w, r)
		return
	}
	// if url is anyting after /categories/ then it will be handled by the root handler
	if r.URL.Path == "/categories/*" {
		handlers.CategoryPostHandler(w, r)
		return
	}
	if r.URL.Path == "/categories" {
		handlers.CategoryHandler(w, r)
		return
	}
	if r.URL.Path == "/creationdates" {
		handlers.CreationDateHandler(w, r)
		return
	}
	if handlers.LikeAllowedURLs(r.URL.Path) {
		handlers.LikeHandler(w, r)
		return
	}
	pages.Error(w, 404)
}
