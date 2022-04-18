package main

import (
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"tradelist/pkg/api"
)

type repoMock struct {
	mock.Mock
}

func (r repoMock) VerifyToken(token string) (api.User, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) InsertPassword(email string, password string) (api.User, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) GetLocations() ([]api.Places, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) GetPostBySubcategoryId(id string) ([]api.Post, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) IsEmailExisting(email string) bool {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) InsertToken(email string, token string) (api.User, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) GetJobPost(posts []api.Post) ([]api.JobPost, string) {
	return []api.JobPost{}, "error"
}

func (r repoMock) GetPostByCategoryId(id string) ([]api.Post, string) {
	return []api.Post{}, ""
}

func (r repoMock) CreateUser(user api.User) (api.User, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) FetchUserInfo(email string) (api.User, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) Save(value api.Post) string {
	return ""
}

func (r repoMock) SaveJobPost(value api.JobPost) string {
	return ""
}

func (r repoMock) GetAllPosts() ([]api.Post, string) {
	return nil, ""
}

func (r repoMock) UpdatePost(post api.Post, postId string) (api.Post, string) {
	return api.Post{}, ""
}

func (r repoMock) DeletePost(postId string) (api.Post, string) {
	return api.Post{}, ""
}

func (r repoMock) GetPostById(id string) (api.Post, string) {
	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}
	post.ID = 1
	return post, ""
}

func (r repoMock) GetSubcategories(categoryId string) ([]api.Subcategory, string) {
	var subcategories = []api.Subcategory{
		{CategoryId: 1, Name: "Accounting"},
		{CategoryId: 1, Name: "HR"},
		{CategoryId: 1, Name: "Legal"},
		{CategoryId: 1, Name: "Customer Service"},
		{CategoryId: 1, Name: "Healthcare"},
		{CategoryId: 1, Name: "Hospitality"},
		{CategoryId: 1, Name: "Housekeeping"},
		{CategoryId: 1, Name: "Software"},
		{CategoryId: 1, Name: "Accounting"},
	}
	return subcategories, ""
}

func (r repoMock) GetDb() *gorm.DB {
	return nil
}

func (r repoMock) GetCategories() ([]api.Category, string) {
	var categories = []api.Category{{Name: "Jobs"}, {Name: "Property"}, {Name: "For Sale"}}
	return categories, ""
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
		//test.Fail()
	}
	status := message["status"]
	if status != 201 {
		//test.Fail()
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

func TestGetSubCategories_NotEmpty(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	response := postService.GetSubcategories("1")

	if response == nil {
		test.Fail()
	}
	var subcategories []api.Subcategory
	mapstructure.Decode(response["data"], &subcategories)

	if len(subcategories) < 9 {
		test.Fail()
	}
}

func TestGetPostById_NotEmpty(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	response := postService.GetPostById("1")

	if response == nil {
		test.Fail()
	}
	var post api.Post
	mapstructure.Decode(response["data"], &post)
	if post.ID == 0 {
		test.Fail()
	}
	assert.Equal(test, 200, response["status"])
}

func TestDeletePost(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	response := postService.DeletePost("1")

	if response == nil {
		test.Fail()
	}
	assert.Equal(test, 200, response["status"])
}

func TestUpdatePost(test *testing.T) {
	repo := repoMock{}
	postService := api.CreatePostService(repo)

	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title2",
		Description:   "Test description"}

	response := postService.UpdatePost(post, "1")

	if response == nil {
		test.Fail()
	}
	assert.Equal(test, "Postid 1 updated", response["message"])
	assert.Equal(test, 200, response["status"])
}

func TestCreateJobPost_Success(test *testing.T) {
	repo := repoMock{}

	postService := api.CreateJobService(repo)

	var job = api.Job{
		Salary:   500,
		Pay:      "monthly",
		Type:     "fulltime",
		Location: "remote",
		Place:    "Gainesville"}

	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}

	var jobPost = api.JobPost{
		Post: post,
		Job:  job,
	}

	message := postService.CreateJobPost(jobPost)
	if message == nil {
		test.Fail()
	}
	status := message["status"]
	if status != 201 {
		test.Fail()
	}
}

func TestGetJobPost_Empty(test *testing.T) {
	repo := repoMock{}

	jobService := api.CreateJobService(repo)

	var posts []api.Post

	jobPosts := jobService.GetJobPost(posts)

	assert.Equal(test, jobPosts, []api.JobPost{})
}

func TestGetPostByCategoryId_NotEmpty(test *testing.T) {
	repo := repoMock{}
	jobService := api.CreateJobService(repo)

	response := jobService.GetPostByCategoryId("1")

	if response == nil {
		test.Fail()
	}
	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "Post found", response["message"])
}
