package main

import (
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"tradelist/mocks"
	"tradelist/pkg/api"
	"tradelist/pkg/data"
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

func GetTestSubcategories(categoryId string) []api.Subcategory {
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
	return subcategories
}

func (r repoMock) GetDb() *gorm.DB {
	return nil
}

func GetTestCategories() []api.Category {
	var categories = []api.Category{{Name: "Jobs"}, {Name: "Property"}, {Name: "For Sale"}}
	return categories
}

func GetTestPost() api.Post {
	return api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title2",
		Description:   "Test description"}
}

func TestCreatePost_Success(test *testing.T) {
	repo := mocks.NewMockRepo(test)

	postService := api.CreatePostService(repo)
	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}

	repo.On("Save", post).Return("")

	response := postService.Create(post)
	assert.Equal(test, 201, response["status"])
	assert.Equal(test, "Post created", response["message"])

}

func TestCreatePost_Error(test *testing.T) {
	repo := mocks.NewMockRepo(test)

	postService := api.CreatePostService(repo)
	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}

	repo.On("Save", post).Return("Error")

	response := postService.Create(post)
	assert.Equal(test, 0, response["status"])
	assert.Equal(test, "Error", response["message"])

}

func TestGetPost_NoRecords(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = []api.Post{}
	repo.On("GetAllPosts").Return(result, "")

	response := postService.GetAllPosts()

	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "No records found", response["message"])
}

func TestGetPost_RecordsFound(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = []api.Post{GetTestPost()}
	repo.On("GetAllPosts").Return(result, "")

	response := postService.GetAllPosts()

	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "Records found", response["message"])
}

func TestGetPost_Error(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = []api.Post{GetTestPost()}
	repo.On("GetAllPosts").Return(result, "error")

	response := postService.GetAllPosts()

	assert.Equal(test, 500, response["status"])
	assert.Equal(test, "error", response["message"])
}

func TestGetCategories_NotEmpty(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = GetTestCategories()
	repo.On("GetCategories").Return(result, "")

	response := postService.GetAllCategories()

	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "Categories found", response["message"])
}

func TestGetCategories_NotFound(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = []api.Category{}
	repo.On("GetCategories").Return(result, "")

	response := postService.GetAllCategories()

	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "No categories found", response["message"])
}

func TestGetCategories_Error(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	var result = []api.Category{}
	repo.On("GetCategories").Return(result, "error")

	response := postService.GetAllCategories()

	assert.Equal(test, 500, response["status"])
	assert.Equal(test, "error", response["message"])
}

func TestGetSubCategories_NotEmpty(test *testing.T) {
	repo := mocks.NewMockRepo(test)
	postService := api.CreatePostService(repo)

	subcategories := data.GetSubcategories()
	repo.On("GetSubcategories", "1").Return(GetSubcategories(subcategories, 1), "")
	repo.On("GetSubcategories", "2").Return(GetSubcategories(subcategories, 2), "")
	repo.On("GetSubcategories", "3").Return(GetSubcategories(subcategories, 3), "")
	repo.On("GetSubcategories", "4").Return(GetSubcategories(subcategories, 4), "")
	repo.On("GetSubcategories", "5").Return(GetSubcategories(subcategories, 5), "")

	response := postService.GetSubcategories("1")

	assert.Equal(test, 200, response["status"])
	assert.Equal(test, "Subcategories found", response["message"])

	var subcategory []api.Subcategory
	mapstructure.Decode(response["data"], &subcategory)
	assert.Equal(test, 9, len(subcategory))

	subcategory = []api.Subcategory{}
	response = postService.GetSubcategories("2")
	mapstructure.Decode(response["data"], &subcategory)
	assert.Equal(test, 5, len(subcategory))

	subcategory = []api.Subcategory{}
	response = postService.GetSubcategories("3")
	mapstructure.Decode(response["data"], &subcategory)
	assert.Equal(test, 8, len(subcategory))

	subcategory = []api.Subcategory{}
	response = postService.GetSubcategories("4")
	mapstructure.Decode(response["data"], &subcategory)
	assert.Equal(test, 8, len(subcategory))

	subcategory = []api.Subcategory{}
	response = postService.GetSubcategories("5")
	mapstructure.Decode(response["data"], &subcategory)
	assert.Equal(test, 4, len(subcategory))
	repo.AssertExpectations(test)
}

func GetSubcategories(subcategories []api.Subcategory, categoryId int) []api.Subcategory {
	var result []api.Subcategory
	for _, item := range subcategories {
		if item.CategoryId == categoryId {
			result = append(result, item)
		}
	}
	print("GetSubcategories ", categoryId, " ", len(result))
	return result
}

//func TestGetPostById_NotEmpty(test *testing.T) {
//	repo := repoMock{}
//	postService := api.CreatePostService(repo)
//
//	response := postService.GetPostById("1")
//
//	if response == nil {
//		test.Fail()
//	}
//	var post api.Post
//	mapstructure.Decode(response["data"], &post)
//	if post.ID == 0 {
//		test.Fail()
//	}
//	assert.Equal(test, 200, response["status"])
//}
//
//func TestDeletePost(test *testing.T) {
//	repo := repoMock{}
//	postService := api.CreatePostService(repo)
//
//	response := postService.DeletePost("1")
//
//	if response == nil {
//		test.Fail()
//	}
//	assert.Equal(test, 200, response["status"])
//}
//
//func TestUpdatePost(test *testing.T) {
//	repo := repoMock{}
//	postService := api.CreatePostService(repo)
//
//	var post = api.Post{
//		SellerId:      1,
//		CategoryId:    1,
//		SubcategoryId: 1,
//		Title:         "Test title2",
//		Description:   "Test description"}
//
//	response := postService.UpdatePost(post, "1")
//
//	if response == nil {
//		test.Fail()
//	}
//	assert.Equal(test, "Postid 1 updated", response["message"])
//	assert.Equal(test, 200, response["status"])
//}
//
//func TestCreateJobPost_Success(test *testing.T) {
//	repo := repoMock{}
//
//	postService := api.CreateJobService(repo)
//
//	var job = api.Job{
//		Salary:   500,
//		Pay:      "monthly",
//		Type:     "fulltime",
//		Location: "remote",
//		Place:    "Gainesville"}
//
//	var post = api.Post{
//		SellerId:      1,
//		CategoryId:    1,
//		SubcategoryId: 1,
//		Title:         "Test title",
//		Description:   "Test description"}
//
//	var jobPost = api.JobPost{
//		Post: post,
//		Job:  job,
//	}
//
//	message := postService.CreateJobPost(jobPost)
//	if message == nil {
//		test.Fail()
//	}
//	status := message["status"]
//	if status != 201 {
//		test.Fail()
//	}
//}
//
//func TestGetJobPost_Empty(test *testing.T) {
//	repo := repoMock{}
//
//	jobService := api.CreateJobService(repo)
//
//	var posts []api.Post
//
//	jobPosts := jobService.GetJobPost(posts)
//
//	assert.Equal(test, jobPosts, []api.JobPost{})
//}
//
//func TestGetPostByCategoryId_NotEmpty(test *testing.T) {
//	repo := repoMock{}
//	jobService := api.CreateJobService(repo)
//
//	response := jobService.GetPostByCategoryId("1")
//
//	if response == nil {
//		test.Fail()
//	}
//	assert.Equal(test, 200, response["status"])
//	assert.Equal(test, "Post found", response["message"])
//}
