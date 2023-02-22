package database

import "forum/server/data"

func GetSpecificPost(posts []data.Post, id int) data.Post {
	if len(posts) == 0 {
		return data.Post{}
	}
	if posts[0].Id == id {
		return posts[0]
	}
	return GetSpecificPost(posts[1:], id)
}
