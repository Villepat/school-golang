package functions

import (
	"fmt"
	"net/http"
)

// Starts the server
func StartServer() {
	fs := http.FileServer(http.Dir("server/templates/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))
	http.HandleFunc("/", StartPage)
	http.HandleFunc("/ascii-art", SubmitData)

	fmt.Println("Server is listening to port #8080...")
	http.ListenAndServe(":8080", nil)
}
