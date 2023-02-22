package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fs := http.FileServer(http.Dir("Front-end/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", StartPage)
	fmt.Println("Server is listening to port #8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
