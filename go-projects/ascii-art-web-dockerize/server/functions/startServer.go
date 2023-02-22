package functions

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fs := http.FileServer(http.Dir("server/templates/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	http.HandleFunc("/", StartPage)
	http.HandleFunc("/ascii-art", SubmitData)

	fmt.Println("Server is listening to port #8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
