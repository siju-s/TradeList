package api

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(user User) (User, string)
	FetchUserInfo(email string) (User, string)
	InsertToken(email string, token string) (User, string)
	VerifyToken(token string) (User, string)
	InsertPassword(email string, password string) (User, string)
	Save(value Post) string
	SaveJobPost(value JobPost) string
	GetJobPost(posts []Post) ([]JobPost, string)
	GetAllPosts() ([]Post, string)
	GetCategories() ([]Category, string)
	GetLocations() ([]Places, string)
	GetSubcategories(categoryId string) ([]Subcategory, string)
	GetPostById(id string) (Post, string)
	GetPostByCategoryId(id string) ([]Post, string)
	GetPostBySubcategoryId(id string) ([]Post, string)
	UpdatePost(post Post, postId string) (Post, string)
	DeletePost(postId string) (Post, string)
	IsEmailExisting(email string) bool
	GetDb() *gorm.DB
}

const (
	AwsUrl = ".s3.amazonaws.com/"
)

type repo struct {
	db *gorm.DB
}

func (r repo) CreateUser(user User) (User, string) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Contact.Password), 14)
	user.Contact.Password = string(password)
	result := r.db.Create(&user)
	return user, handleError(result.Error)
}

func (r repo) FetchUserInfo(email string) (User, string) {
	var user User
	result := r.db.Where("email= ?", email).First(&user)
	return user, handleError(result.Error)
}

func (r repo) InsertToken(email string, token string) (User, string) {
	var user User
	result := r.db.Model(&user).Where("email= ?", email).Update("Token", token)
	return user, handleError(result.Error)
}

func (r repo) VerifyToken(token string) (User, string) {
	var user User
	result := r.db.Where("Token= ?", token).First(&user)
	return user, handleError(result.Error)
}

func (r repo) InsertPassword(email string, password string) (User, string) {
	var user User
	result := r.db.Model(&user).Where("email= ?", email).Update("password", password)
	return user, handleError(result.Error)
}

func (r repo) GetPostById(id string) (Post, string) {
	var post Post
	err := r.db.First(post, id).Find(&id).Error
	return post, handleError(err)
}

func (r repo) GetPostByCategoryId(id string) ([]Post, string) {
	var post []Post
	err := r.db.Where("category_id = ?", id).Order("created_at desc").Find(&post).Error
	var bucketid = GetEnvWithKey("AWS_BUCKET")
	if err == nil {
		fetchImages(post, r, bucketid)
	}
	return post, handleError(err)
}

func (r repo) GetPostBySubcategoryId(id string) ([]Post, string) {
	var post []Post
	err := r.db.Where("subcategory_id = ?", id).Order("created_at desc").Find(&post).Error
	var bucketid = GetEnvWithKey("AWS_BUCKET")
	if err == nil {
		fetchImages(post, r, bucketid)
	}
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
	err := r.db.Where("category_id = ?", categoryId).Find(&subcategories).Error
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

func (r repo) GetLocations() ([]Places, string) {
	var places []Places
	err := r.db.Find(&places).Error
	return places, handleError(err)
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

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func (r repo) GetAllPosts() ([]Post, string) {
	var bucketid = GetEnvWithKey("AWS_BUCKET")
	var posts []Post
	err := r.db.Order("created_at desc").Find(&posts).Error
	if err == nil {
		fetchImages(posts, r, bucketid)
	}
	return posts, handleError(err)
}

func fetchImages(posts []Post, r repo, bucketid string) {
	for idx, post := range posts {
		images, err := r.GetImagesForPost(post.ID)
		fmt.Println("Post id:", post.ID, "image", len(images), "idx:", idx)
		if err == "" {
			for i, img := range images {
				images[i].Url = "https://" + bucketid + AwsUrl + img.Url
			}
			posts[idx].Image = images
		}
	}
}

func (r repo) GetImagesForPost(postId uint) ([]Images, string) {
	var images []Images
	err := r.db.Where("post_id = ?", postId).Find(&images).Error
	return images, handleError(err)
}

func (r repo) Save(value Post) string {
	result := r.db.Save(&value).Error
	fmt.Println(result)
	if result == nil {
		return ""
	}
	return result.Error()
}

func (r repo) SaveJobPost(jobPost JobPost) string {
	result := r.db.Save(&jobPost.Post).Error
	fmt.Println(result)
	if result == nil {
		jobPost.Job.PostId = jobPost.Post.ID
		jobPost.Job.SubcategoryId = jobPost.Post.SubcategoryId
		r.db.Save(&jobPost.Job)
		return ""
	}
	return result.Error()
}

func (r repo) GetJobPost(posts []Post) ([]JobPost, string) {
	var jobPosts []JobPost
	var e string
	for _, post := range posts {
		var job Job
		err := r.db.Where("post_id = ?", post.ID).Find(&job).Error
		if err == nil {
			var jobPost JobPost
			jobPost.Post = post
			jobPost.Job = job

			jobPosts = append(jobPosts, jobPost)
		} else {
			e = err.Error()
		}
	}
	return jobPosts, e
}

func (r repo) IsEmailExisting(email string) bool {
	var user User
	rows := r.db.Where("email = ?", email).Find(&user).RowsAffected
	return rows > 0
}

func CreateRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
