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
	GetLocations() map[string]interface{}
	GetSubcategories(categoryId string) map[string]interface{}
	GetPostById(categoryId string) map[string]interface{}
	UpdatePost(post Post, postId string) map[string]interface{}
	DeletePost(postId string) map[string]interface{}
	GetPostsByUser(id string) map[string]interface{}
	//GetDb() *gorm.DB
	GetRepo() Repo
}

type postService struct {
	db   *gorm.DB
	repo Repo
}

func (service *postService) GetRepo() Repo {
	return service.repo
}

//func CreatePostService(db *gorm.DB) PostService {
//	return &postService{db: db}
//}

func CreatePostService(repo Repo) PostService {
	return &postService{repo: repo}
}

//func (service *postService) GetDb() *gorm.DB {
//	return service.db
//}

func (service *postService) Create(post Post) map[string]interface{} {
	err := service.repo.Save(post)

	if err != "" {
		return apihelpers.Message(0, err)
	} else {
		return apihelpers.Message(http.StatusCreated, "Post created")
	}
}

func (service *postService) GetAllPosts() map[string]interface{} {
	posts, err := service.repo.GetAllPosts()
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
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

func (service *postService) GetPostsByUser(id string) map[string]interface{} {
	posts, err := service.repo.GetPostsByUser(id)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
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
	categories, err := service.repo.GetCategories()
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
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

func (service *postService) GetLocations() map[string]interface{} {
	places, err := service.repo.GetLocations()
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	size := len(places)

	var message string
	if size > 0 {
		message = "Locations found"
	} else {
		message = "No locations found"
	}
	response := apihelpers.Message(http.StatusOK, message)
	response["data"] = places
	return response
}

func (service *postService) GetSubcategories(categoryId string) map[string]interface{} {
	subcategories, err := service.repo.GetSubcategories(categoryId)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
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
	post, err := service.repo.GetPostById(id)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	response["data"] = post
	return response
}

func (service *postService) UpdatePost(post Post, postId string) map[string]interface{} {
	_, err := service.repo.UpdatePost(post, postId)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	response := apihelpers.Message(http.StatusOK, "Postid "+postId+" updated")
	return response
}

func (service *postService) DeletePost(postId string) map[string]interface{} {
	_, err := service.repo.DeletePost(postId)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	response := apihelpers.Message(http.StatusOK, "Postid "+postId+" deleted")
	return response
}
