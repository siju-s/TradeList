package api

import "gorm.io/gorm"

type Repo interface {
	Save(value interface{}) string
	GetAllPosts() ([]Post, string)
	GetCategories() ([]Category, string)
	GetDb() *gorm.DB
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetDb() *gorm.DB {
	return r.db
}

func (r repo) GetCategories() ([]Category, string) {
	var categories []Category
	err := r.db.Find(&categories).Error
	return categories, handleError(err)
}

func handleError(err error) string {
	var msg string
	if err == nil {
		msg = ""
	} else {
		msg = err.Error()
	}
	return msg
}

func (r repo) GetAllPosts() ([]Post, string) {
	var posts []Post
	err := r.db.Find(&posts).Error
	return posts, handleError(err)
}

func (r repo) Save(value interface{}) string {
	return r.db.Save(&value).Error.Error()
}

func CreateRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
