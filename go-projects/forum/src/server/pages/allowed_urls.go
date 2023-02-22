package pages

import (
	"strconv"

	"forum/server/database"
)

// AllowedURLs returns a list of all the allowed URLs for posts
func AllowedURLs() []string {
	s := ""
	posts, _ := database.GetPosts(s)
	results := []string{}
	for _, v := range posts {
		results = append(results, "/"+v.Author+"/"+strconv.Itoa(v.Id)+"/"+database.MakeURL(v.Title, ""))
	}
	return results
}
