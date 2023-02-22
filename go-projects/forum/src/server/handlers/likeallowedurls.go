package handlers

import (
	"strconv"
	"strings"
)

func LikeAllowedURLs(url string) bool {
	array := strings.Split(url, "/")
	if len(array) != 4 {
		return false
	}
	if array[1] != "liked" && array[1] != "disliked" {
		return false
	}
	if _, err := strconv.Atoi(array[2]); err != nil {
		return false
	}
	if array[3] != "comment" && array[3] != "post" {
		return false
	}
	return true
}
