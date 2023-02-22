package handlers

import (
	"encoding/json"
	"fmt"
	"forum/server/database"
	"net/http"
)

func CreationDateHandler(w http.ResponseWriter, r *http.Request) {
	posts := database.GetCreationDates(database.Db)
	fmt.Println(posts)
	json.NewEncoder(w).Encode(posts)
}
