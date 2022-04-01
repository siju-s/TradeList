package api

import (
	"fmt"
	"net/http"
	"tradelist/pkg/apihelpers"

	"gorm.io/gorm"
)

type LoginService interface {
	SignUp(user User) map[string]interface{}
	LoginEmail(user User, email string) map[string]interface{}
}

type loginService struct {
	db   *gorm.DB
	repo Repo
}

func CreateService(repo Repo) LoginService {
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

func (service loginService) LoginEmail(user User, email string) map[string]interface{} {
	user, err := service.repo.LoginEmail(user, email)
	fmt.Println("User id:", user.ID)
	if err != "" {
		return apihelpers.Message(0, err)
	} else {
		return apihelpers.Message(http.StatusCreated, "User exists")
	}
}
