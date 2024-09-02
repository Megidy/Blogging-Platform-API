package models

import (
	"github.com/Megidy/BloggingPlatformApi/pkj/config"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:""json:"title"`
	Content  string `json:"author"`
	Category string `json:"category"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (post *Post) CreatePost() *Post {
	db.Create(post)
	return post
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Find(Posts)
	return Posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db.Where("ID=?", Id).Find(getPost)
	return &getPost, db
}

func DeletePost(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}
