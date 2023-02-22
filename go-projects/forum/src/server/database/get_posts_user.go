package database

import (
	"forum/server/data"
	"log"
)

func GetPostsUser(user string) ([]data.Post, error) {
	result := []data.Post{}
	var err error
	Rows, err = Db.Query("SELECT * FROM posts WHERE Author=? ORDER BY id DESC", user)
	if err != nil {
		log.Println(err)
		return result, err
	}
	result = userscan()
	for i, v := range result {
		result[i].URL = MakeURL(v.Title, "")
	}
	return result, err
}
