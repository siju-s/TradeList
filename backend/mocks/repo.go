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

func (r repoMock) GetPostsByUser(id string) ([]api.Result, string) {
	//TODO implement me
	panic("implement me")
}

func NewMockRepo(t *testing.T) *repoMock {
	t.Helper()
	return &repoMock{}
}

func (r repoMock) IsEmailExisting(email string) bool {
	args := r.Called(email)
	return args.Get(0).(bool)
}

func (r repoMock) VerifyToken(token string) (api.User, string) {

	args := r.Called(token)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) InsertPassword(email string, password string) (api.User, string) {

	args := r.Called(email, password)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) GetLocations() ([]api.Places, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) GetPostBySubcategoryId(id string) ([]api.Result, string) {
	//TODO implement me
	panic("implement me")
}

func (r repoMock) InsertToken(email string, token string) (api.User, string) {

	args := r.Called(email, token)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) GetJobPost(posts []api.Post) ([]api.JobPost, string) {
	return []api.JobPost{}, "error"
}

func (r repoMock) GetPostByCategoryId(id string) ([]api.Result, string) {
	return []api.Result{}, ""
}

func (r repoMock) CreateUser(user api.User) (api.User, string) {

	args := r.Called(user)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) FetchUserInfo(email string) (api.User, string) {
	args := r.Called(email)
	return args.Get(0).(api.User), args.Get(1).(string)
}

func (r repoMock) Save(value api.Post) string {
	args := r.Called(value)
	return args.Get(0).(string)
}

func (r repoMock) SaveJobPost(value api.JobPost) string {
	return ""
}

func (r repoMock) GetAllPosts() ([]api.Result, string) {
	args := r.Called()
	return args.Get(0).([]api.Result), args.Get(1).(string)
}

func (r repoMock) UpdatePost(post api.Post, postId string, userId string) (api.Post, string, int64) {
	return api.Post{}, "", 0
}

func (r repoMock) DeletePost(postId string, userId string) (api.Post, string) {
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
	args := r.Called(categoryId)
	return args.Get(0).([]api.Subcategory), args.Get(1).(string)
}

func (r repoMock) GetDb() *gorm.DB {
	return nil
}

func (r repoMock) GetCategories() ([]api.Category, string) {
	args := r.Called()
	return args.Get(0).([]api.Category), args.Get(1).(string)
}
