package pages

import (
	"forum/server/database"
)

func AllowedCategoriesURLs() []string {
	results := []string{}
	tags, _ := database.GetTags(0)
	for _, v := range tags {
		results = append(results, "/categories/"+database.MakeURL(v, ""))
	}
	return results
}
