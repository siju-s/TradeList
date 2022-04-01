package api

import (
	"net/http"
	"tradelist/pkg/apihelpers"
)

type JobService interface {
	CreateJobPost(jobPost JobPost) map[string]interface{}
	GetJobPost(posts []Post) []JobPost
	GetPostByCategoryId(id string) map[string]interface{}
}

type jobService struct {
	repo Repo
}

func (service *jobService) GetRepo() Repo {
	return service.repo
}

func CreateJobService(repo Repo) JobService {
	return &jobService{repo: repo}
}

func (service *jobService) CreateJobPost(jobPost JobPost) map[string]interface{} {
	err := service.repo.SaveJobPost(jobPost)

	if err != "" {
		return apihelpers.Message(0, err)
	} else {
		return apihelpers.Message(http.StatusCreated, "Post created")
	}
}

func (service *jobService) GetPostByCategoryId(id string) map[string]interface{} {
	var posts []Post
	posts, err := service.repo.GetPostByCategoryId(id)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	if id == "1" && len(posts) > 0 {
		response["data"] = service.GetJobPost(posts)
	}
	return response
}

func (service *jobService) GetJobPost(posts []Post) []JobPost {
	jobPosts, err := service.repo.GetJobPost(posts)
	if err != "" {
		return []JobPost{}
	}
	return jobPosts
}
