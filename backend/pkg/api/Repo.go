package api

import (
	"fmt"

	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(user User) (User, string)
	FetchUserInfo(email string) (User, string)
	Save(value Post) string
	SaveJobPost(value JobPost) string
	GetAllPosts(bucketid string) ([]Post, string)
	GetCategories() ([]Category, string)
	GetSubcategories(categoryId string) ([]Subcategory, string)
	GetPostById(id string) (Post, string)
	UpdatePost(post Post, postId string) (Post, string)
	DeletePost(postId string) (Post, string)
	GetDb() *gorm.DB
}

const (
	AwsUrl = ".s3.amazonaws.com/"
)

type repo struct {
	db *gorm.DB
}

func (r repo) CreateUser(user User) (User, string) {
	result := r.db.Create(&user)
	return user, handleError(result.Error)
}

func (r repo) FetchUserInfo(email string) (User, string) {
	var user User
	result := r.db.Where("email= ?", email).First(&user)
	return user, handleError(result.Error)
}

func (r repo) GetPostById(id string) (Post, string) {
	var post Post
	err := r.db.First(post, id).Find(&id).Error
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
	err := r.db.Find(&subcategories, categoryId).Error
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

func handleError(err error) string {
	var msg string
	if err == nil {
		msg = ""
	} else {
		msg = err.Error()
	}
	return msg
}

func (r repo) GetAllPosts(bucketid string) ([]Post, string) {
	var posts []Post
	err := r.db.Find(&posts).Error
	if err == nil {
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
	return posts, handleError(err)
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

func CreateRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
