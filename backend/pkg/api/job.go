package api

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"tradelist/pkg/apihelpers"
)

type JobService interface {
	CreateJobPost(jobPost JobPost) map[string]interface{}
	GetJobPost(posts []Post) []JobPost
	GetPostByCategoryId(id string) map[string]interface{}
}

type jobService struct {
	db *gorm.DB
}

func CreateJobService(db *gorm.DB) JobService {
	return &jobService{db: db}
}

func (service *jobService) CreateJobPost(jobPost JobPost) map[string]interface{} {
	err := service.db.Debug().Save(&jobPost.Post).Error
	fmt.Println("Postid:", jobPost.Post.ID)
	if err == nil {
		jobPost.Job.PostId = jobPost.Post.ID
		jobPost.Job.SubcategoryId = jobPost.Post.SubcategoryId
		err = service.db.Debug().Save(&jobPost.Job).Error
	}
	return apihelpers.Message(http.StatusCreated, "JobPost created")
}

func (service *jobService) GetPostByCategoryId(id string) map[string]interface{} {
	var posts []Post
	err := service.db.Debug().Where("category_id = ?", id).Find(&posts).Error
	if err != nil {
		return apihelpers.Message(http.StatusInternalServerError, err.Error())
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	if id == "1" {
		response["data"] = service.GetJobPost(posts)
	}
	return response
}

func (service *jobService) GetJobPost(posts []Post) []JobPost {
	var jobPosts []JobPost
	for _, post := range posts {
		var job Job
		err := service.db.Debug().Where("post_id = ?", post.ID).Find(&job).Error
		if err == nil {
			var jobPost JobPost
			jobPost.Post = post
			jobPost.Job = job

			jobPosts = append(jobPosts, jobPost)
		}
	}
	return jobPosts
}
