package api

import (
	"fmt"
	"net/http"
	"tradelist/pkg/apihelpers"

	"gorm.io/gorm"
)

type LoginService interface {
	SignUp(user User) map[string]interface{}
	FetchUserInfo(email string) (User, map[string]interface{})
}

type loginService struct {
	db   *gorm.DB
	repo Repo
}

func CreateLoginService(repo Repo) LoginService {
	return &loginService{repo: repo}
}

func (service loginService) SignUp(user User) map[string]interface{} {
	user, err := service.repo.CreateUser(user)
	fmt.Println("User id:", user.ID)
	if err != "" {
		return apihelpers.Message(0, err)
	} else {
		return apihelpers.Message(http.StatusCreated, "User created successfully")
	}
}

func (service loginService) FetchUserInfo(email string) (User, map[string]interface{}) {
	user, err := service.repo.FetchUserInfo(email)
	fmt.Println("User id:", user.ID)
	var response map[string]interface{}
	if err != "" {
		user.ID = 0
		response = apihelpers.Message(http.StatusNotFound, "User not found")
	}
	return user, response
}
