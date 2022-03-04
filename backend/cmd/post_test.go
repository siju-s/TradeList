package main

import (
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"tradelist/pkg/api"
)

type repoMock struct {
	mock.Mock
}

func (r repoMock) GetDb() *gorm.DB {
	return nil
}

func (r repoMock) GetCategories() ([]api.Category, string) {
	var categories = []api.Category{{Name: "Jobs"}, {Name: "Property"}, {Name: "For Sale"}}
	return categories, ""
}

func (r repoMock) GetAllPosts() ([]api.Post, string) {
	return nil, ""
}

func (r repoMock) Save(value interface{}) string {
	return ""
}

func TestCreatePost_Success(test *testing.T) {
	repo := repoMock{}

	postService := api.CreatePostService(repo)
	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}

	message := postService.Create(post)
	if message == nil {
		test.Fail()
	}
	status := message["status"]
	if status != 201 {
		test.Fail()
	}
}

func TestGetPost_Empty(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	message := postService.GetAllPosts()

	if message == nil {
		test.Fail()
	}
	msg := message["message"]

	if msg != "No records found" {
		test.Fail()
	}
}

func TestGetCategories_NotEmpty(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	response := postService.GetAllCategories()

	if response == nil {
		test.Fail()
	}
	var categories []api.Category
	mapstructure.Decode(response["data"], &categories)
	if len(categories) < 3 {
		test.Fail()
	}
}
