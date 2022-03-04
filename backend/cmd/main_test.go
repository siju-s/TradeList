package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http/httptest"
	"os"
	"testing"
	"tradelist/pkg/api"
	"tradelist/pkg/app"
)

var server *app.Server

const NUM_CATEGORIES = 3

func TestMain(m *testing.M) {
	provideMock()
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func provideMock() *app.Server {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil || db == nil {
		panic(err.Error())
	}
	repo := api.CreateRepo(db)
	fmt.Println("Mock called")

	postService := api.CreatePostService(repo)
	jobService := api.CreateJobService(db)
	server = app.CreateServer(mux.NewRouter(), postService, jobService)
	server.Migrate(db)
	server.RunTest()
	return server
}

func TestGetCategories(test *testing.T) {
	response := sendRequest(test, "GET", "/categories", nil)
	var categories []api.Category
	mapstructure.Decode(response["data"], &categories)
	if len(categories) != NUM_CATEGORIES {
		test.Errorf("Categories cannot be empty")
	}
}

func TestGetPostsReturnsEmpty(test *testing.T) {
	deletePosts()
	response := sendRequest(test, "GET", "/post", nil)
	var posts []api.Post
	mapstructure.Decode(response["data"], &posts)
	if len(posts) != 0 {
		test.Errorf("Invalid posts length, expected 0")
	}
}

func TestCreatePost(test *testing.T) {
	response := createPost(test)

	if len(response) == 0 || response["message"] != "Post created" {
		test.Fail()
	}
	fmt.Println(response)
}

func createPost(test *testing.T) map[string]interface{} {
	var post = api.Post{
		SellerId:      1,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}
	body, _ := json.Marshal(post)
	response := sendRequest(test, "POST", "/post", bytes.NewReader(body))
	return response
}

func TestCreatePost_Fail_Constraint(test *testing.T) {
	var post = api.Post{
		SellerId:      5,
		CategoryId:    1,
		SubcategoryId: 1,
		Title:         "Test title",
		Description:   "Test description"}
	body, _ := json.Marshal(post)
	response := sendRequest(test, "POST", "/post", bytes.NewReader(body))

	if len(response) == 0 || response["message"] != "FOREIGN KEY constraint failed" {
		test.Fail()
	}
	fmt.Println(response)
}

func TestGetPosts_AfterCreate_Returns_Record(test *testing.T) {
	deletePosts()
	createPost(test)
	response := sendRequest(test, "GET", "/post", nil)
	var posts []api.Post
	mapstructure.Decode(response["data"], &posts)
	if len(posts) == 0 {
		test.Errorf("Expected Post data but no records found")
	}
}

func sendRequest(test *testing.T, method string, endpoint string, body io.Reader) map[string]interface{} {
	ts := httptest.NewServer(server.Router)
	defer ts.Close()
	req := httptest.NewRequest(method, endpoint, body)
	rr := httptest.NewRecorder()

	server.Router.ServeHTTP(rr, req)

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		test.Errorf(" Error getting data")
	}
	return response
}

func deletePosts() {
	db := server.PostService.GetRepo().GetDb()
	session := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped()
	session.Delete(&api.Post{})
}

func cleanup() {
	db := server.PostService.GetRepo().GetDb()
	session := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped()
	session.Delete(&api.Subcategory{})
	session.Delete(&api.Category{})
	session.Delete(&api.Post{})
	session.Delete(&api.Seller{})
}
