package data

import (
	"html/template"
)

type Post struct {
	Id           int
	Author       string
	Title        string
	Content      string
	CreationDate string
	Comments     []Comment
	Len_Comments int
	LoggedIn     bool
	Views        int
	Tags         string
	URL          string
	LikeStatus   int
	Likes        int
	Dislikes     int
}

type Comment struct {
	Id           int
	Author       string
	Content      string
	CreationDate string
	LikeStatus   int
	Username     string
	LoggedIn     bool
	Likes        int
	Dislikes     int
}

type Like struct {
	Id        int
	Status    int
	Username  string
	CommentId int
	PostId    int
}

type RegisterStruct struct {
	User     string
	Password string
	Email    string
	Navbar   template.HTML
}

type CreationDates struct {
	PostID   int
	Creation string
}

type MainStruct struct {
	LoggedIn bool
	Username string
	Data     interface{}
	Tags     []string
}

type HomeStruct struct {
	Data     string
	Topics   []Post
	LoggedIn bool
}

type CategoryStruct struct {
	Tags   []string
	TagsJS string
}

type ErrorStruct struct {
	Data string
}

type LoginStruct struct {
	Data string
}

type User struct {
	Name           string
	ProfilePicture string
	Posts          []Post
}

type RegisteredUser struct {
	TheUser  User
	LoggedIn bool
}

type UserProfile struct {
	Data         []string
	Posts        []Post
	URLs         []string
	Likes        []Post
	CommentLikes []Post
	//LoggedIn bool
}
