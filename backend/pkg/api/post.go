package api

import (
	"gorm.io/gorm"
	"net/http"
	"tradelist/pkg/apihelpers"
)

type PostService interface {
	Create(post Post) map[string]interface{}
	GetAllPosts() map[string]interface{}
	GetAllCategories() map[string]interface{}
	GetSubcategories(categoryId string) map[string]interface{}
	GetPostById(categoryId string) map[string]interface{}
	UpdatePost(post Post, postId string) map[string]interface{}
	DeletePost(postId string) map[string]interface{}
}

type postService struct {
	db *gorm.DB
}

func CreatePostService(db *gorm.DB) PostService {
	return &postService{db: db}
}

func (service *postService) Create(post Post) map[string]interface{} {
	err := service.db.Save(&post).Error
	if err != nil {
		return apihelpers.Message(0, err.Error())
	} else {
		return apihelpers.Message(http.StatusCreated, "Post created")
	}
}

func (service *postService) GetAllPosts() map[string]interface{} {
	var posts []Post
	err := service.db.Find(&posts).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	} else {
		size := len(posts)

		var message string
		if size > 0 {
			message = "Records found"
		} else {
			message = "No records found"
		}
		response := apihelpers.Message(http.StatusOK, message)
		response["data"] = posts
		return response
	}
}

func (service *postService) GetAllCategories() map[string]interface{} {
	var categories []Category
	err := service.db.Find(&categories).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	size := len(categories)

	var message string
	if size > 0 {
		message = "Categories found"
	} else {
		message = "No categories found"
	}
	response := apihelpers.Message(http.StatusOK, message)
	response["data"] = categories
	return response
}

func (service *postService) GetSubcategories(categoryId string) map[string]interface{} {
	var subcategories []Subcategory
	err := service.db.Find(&subcategories, categoryId).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	size := len(subcategories)

	var message string
	if size > 0 {
		message = "Subcategories found"
	} else {
		message = "No subcategories found"
	}
	response := apihelpers.Message(http.StatusOK, message)
	response["data"] = subcategories
	return response
}

func (service *postService) GetPostById(id string) map[string]interface{} {
	var post Post
	err := service.db.First(post, id).Find(&id).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	response["data"] = post
	return response
}

func (service *postService) UpdatePost(post Post, postId string) map[string]interface{} {
	var postData Post
	err := service.db.First(&postData, postId).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}

	err = service.db.Where("ID = ?", postId).Updates(&post).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	response := apihelpers.Message(http.StatusOK, "Postid "+postId+" updated")
	return response
}

func (service *postService) DeletePost(postId string) map[string]interface{} {
	var post Post
	err := service.db.Delete(&post, postId).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	response := apihelpers.Message(http.StatusOK, "Postid "+postId+" deleted")
	return response
}
