package api

import (
	"net/http"
	"tradelist/pkg/apihelpers"
)

type JobService interface {
	CreateJobPost(jobPost JobPost) map[string]interface{}
	GetJobPost(posts []Post) []JobPost
	GetPostByCategoryId(id string) map[string]interface{}
	GetPostBySubcategoryId(id string) map[string]interface{}
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
	var results []Result
	results, err := service.repo.GetPostByCategoryId(id)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	if id == "1" && len(results) > 0 {
		var posts []Post
		for i := range results {
			posts = append(posts, results[i].Post)
		}
		jobPosts := service.GetJobPost(posts)
		for i := range results {
			results[i].Job = jobPosts[i].Job
		}
		response["data"] = results
	}
	return response
}

func (service *jobService) GetPostBySubcategoryId(id string) map[string]interface{} {
	var results []Result
	results, err := service.repo.GetPostBySubcategoryId(id)
	if err != "" {
		return apihelpers.Message(http.StatusInternalServerError, err)
	} else if len(results) == 0 {
		return apihelpers.Message(http.StatusOK, "No posts found")
	}
	response := apihelpers.Message(http.StatusOK, "Post found")
	if id == "1" {
		var posts []Post
		for i := range results {
			posts = append(posts, results[i].Post)
		}
		jobPosts := service.GetJobPost(posts)
		for i := range results {
			results[i].Job = jobPosts[i].Job
		}
		response["data"] = results
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
