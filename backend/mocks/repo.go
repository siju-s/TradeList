package mocks

import (
	"testing"
	"tradelist/pkg/api"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type repoMock struct {
	mock.Mock
}

func NewMockRepo(t *testing.T) *repoMock {
	t.Helper()
	return &repoMock{}
}

func (r repoMock) IsEmailExisting(email string) bool {
	args := r.Called(email)
	//if args.Error(0) != nil {
	//	return false
	//}
	return args.Get(0).(bool)
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
	args := r.Called(user)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) FetchUserInfo(email string) (api.User, string) {
	args := r.Called(email)
	//if args.Error(0) != nil {
	//	return false
	//}
	return args.Get(0).(api.User), args.Get(1).(string)
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
