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
	GetAllPosts() ([]Result, string)
	GetCategories() ([]Category, string)
	GetLocations() ([]Places, string)
	GetSubcategories(categoryId string) ([]Subcategory, string)
	GetPostById(id string) (Post, string)
	GetPostByCategoryId(id string) ([]Result, string)
	GetPostBySubcategoryId(id string) ([]Result, string)
	UpdatePost(post Post, postId string, userId string) (Post, string, int64)
	DeletePost(postId string, userid string) (Post, string)
	IsEmailExisting(email string) bool
	GetPostsByUser(id string) ([]Result, string)
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
	err := r.db.First(&post, id).Error
	return post, handleError(err)
}

func (r repo) GetPostByCategoryId(id string) ([]Result, string) {
	var results []Result
	err := r.db.Table("posts").Select("posts.*, users.id, first_name, last_name, email, phone").
		Joins("join users on users.id = posts.seller_id").
		Where("category_id = ?", id).
		Order("created_at desc").
		Scan(&results).
		Error
	var bucketId = GetEnvWithKey("AWS_BUCKET")
	if err == nil {
		fetchImages(results, r, bucketId)
	}
	return results, handleError(err)
}

func (r repo) GetPostsByUser(id string) ([]Result, string) {
	var results []Result
	err := r.db.Table("posts").Select("posts.*, users.id, first_name, last_name, email, phone").
		Joins("join users on users.id = posts.seller_id").
		Where("seller_id = ?", id).
		Order("created_at desc").
		Scan(&results).
		Error
	var bucketId = GetEnvWithKey("AWS_BUCKET")
	if err == nil {
		fetchImages(results, r, bucketId)
	}
	return results, handleError(err)
}

func (r repo) GetPostBySubcategoryId(id string) ([]Result, string) {
	var results []Result
	err := r.db.Table("posts").Select("posts.*, users.id, first_name, last_name, email, phone").
		Joins("join users on users.id = posts.seller_id").
		Where("subcategory_id = ?", id).
		Order("created_at desc").
		Scan(&results).
		Error
	var bucketid = GetEnvWithKey("AWS_BUCKET")
	if err == nil {
		fetchImages(results, r, bucketid)
	}
	return results, handleError(err)
}

func (r repo) UpdatePost(post Post, postId string, userId string) (Post, string, int64) {
	var postData Post
	result := r.db.Where("ID = ? and seller_id = ?", postId, userId).Updates(&postData)
	return post, handleError(result.Error), result.RowsAffected
}

func (r repo) DeletePost(postId string, userid string) (Post, string) {
	var post Post
	err := r.db.Where("seller_id = ? and id = ?", userid, postId).Delete(&post).Error
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

type Result struct {
	Post Post `gorm:"embedded"`
	User User `gorm:"embedded"`
	Job  Job  `gorm:"embedded"`
}

func (r repo) GetAllPosts() ([]Result, string) {
	var bucketid = GetEnvWithKey("AWS_BUCKET")
	var results []Result
	err := r.db.Table("posts").Select("posts.*, users.id, first_name, last_name, email, phone").
		Joins("join users on users.id = posts.seller_id").
		Order("created_at desc").
		Scan(&results).
		Error

	if err == nil {
		// FIXME Temporary solution as join with image table not working
		fetchImages(results, r, bucketid)
	}
	return results, handleError(err)
}

func fetchImages(posts []Result, r repo, bucketid string) {
	for idx, post := range posts {
		var postData = post.Post
		images, err := r.GetImagesForPost(postData.ID)
		fmt.Println("Post id:", postData.ID, "image", len(images), "idx:", idx)
		if err == "" {
			for i, img := range images {
				images[i].Url = "https://" + bucketid + AwsUrl + img.Url
			}
			posts[idx].Post.Image = images
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
