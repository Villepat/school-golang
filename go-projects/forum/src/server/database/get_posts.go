package database

import (
	"forum/server/data"
	"log"
)

func GetPosts(tag string) ([]data.Post, error) {
	result := []data.Post{}
	var err error
	if tag == "" {
		Rows, err = Db.Query("SELECT * FROM posts ORDER BY id DESC")
		if err != nil {
			log.Println(err)
			return result, err
		}
		result = userscan()

	} else {
		// find posts with tag or "*,tag" or "tag,*" or "*,tag,*"
		Rows, err = Db.Query("SELECT * FROM posts WHERE tags LIKE ? OR tags LIKE ? OR tags LIKE ? OR tags LIKE ? ORDER BY id DESC", "%,"+tag, tag+",%", "%,"+tag+",%", tag)
		if err != nil {
			log.Panic(err)
			return result, err
		}
		result = userscan()
	}
	for i, v := range result {
		result[i].URL = MakeURL(v.Title, "")
		result[i].Len_Comments = len(GetComments(v.Id))
	}
	return result, err
}

func userscan() []data.Post {
	result := []data.Post{}
	var err error
	defer Rows.Close()
	for Rows.Next() {
		var id int
		var Author string
		var Title string
		var Content string
		var Created string
		var Views int
		var Tags string
		err = Rows.Scan(&id, &Author, &Title, &Content, &Created, &Views, &Tags)
		if err != nil {
			log.Panic(err)
		}
		post := data.Post{
			Id:           id,
			Author:       Author,
			Title:        Title,
			Content:      Content,
			CreationDate: Created,
			Views:        Views,
			Tags:         Tags,
		}
		result = append(result, post)
	}
	return result
}
