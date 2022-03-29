package api

import (
	"fmt"
	"gorm.io/gorm"
)

type Repo interface {
	Save(value Post) string
	GetAllPosts() ([]Post, string)
	GetCategories() ([]Category, string)
	GetSubcategories(categoryId string) ([]Subcategory, string)
	GetPostById(id string) (Post, string)
	UpdatePost(post Post, postId string) (Post, string)
	DeletePost(postId string) (Post, string)
	GetDb() *gorm.DB
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetPostById(id string) (Post, string) {
	var post Post
	err := r.db.First(post, id).Find(&id).Error
	return post, handleError(err)
}

func (r repo) UpdatePost(post Post, postId string) (Post, string) {
	var postData Post
	err := r.db.First(&postData, postId).Error

	err = r.db.Where("ID = ?", postId).Updates(&post).Error
	return post, handleError(err)
}

func (r repo) DeletePost(postId string) (Post, string) {
	var post Post
	err := r.db.Delete(&post, postId).Error

	return post, handleError(err)
}

func (r repo) GetSubcategories(categoryId string) ([]Subcategory, string) {
	var subcategories []Subcategory
	err := r.db.Find(&subcategories, categoryId).Error
	return subcategories, handleError(err)
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

func (r repo) Save(value Post) string {
	result := r.db.Save(&value).Error
	fmt.Println(result)
	if result == nil {
		return ""
	}
	return result.Error()
}

func CreateRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
