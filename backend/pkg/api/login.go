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
	InsertToken(email string, token string) (User, map[string]interface{})
	VerifyToken(token string) (User, map[string]interface{})
	InsertPassword(email string, password string) (User, map[string]interface{})
}

type loginService struct {
	db   *gorm.DB
	repo Repo
}

func CreateLoginService(repo Repo) LoginService {
	return &loginService{repo: repo}
}

func (service loginService) FetchUserInfo(email string) (User, map[string]interface{}) {
	user, err := service.repo.FetchUserInfo(email)
	fmt.Println("User id:", user.ID)
	var response map[string]interface{}
	if err != "" {
		user.ID = 0
		response = apihelpers.Message(http.StatusNotFound, "User not found")
		return user, response
	}
	response = apihelpers.Message(http.StatusCreated, "User found")

	return user, response
}

func (service loginService) InsertToken(email string, token string) (User, map[string]interface{}) {
	user, err := service.repo.InsertToken(email, token)
	fmt.Println("User id:", user.ID)
	var response map[string]interface{}
	if err != "" {
		user.ID = 0
		response = apihelpers.Message(http.StatusNotFound, "User not found")
	}
	return user, response
}

func (service loginService) VerifyToken(token string) (User, map[string]interface{}) {
	user, err := service.repo.VerifyToken(token)
	fmt.Println("User id:", user.ID)
	var response map[string]interface{}
	if err != "" {
		user.ID = 0
		response = apihelpers.Message(http.StatusNotFound, "Invalid Token")
	}
	return user, response
}

func (service loginService) InsertPassword(email string, password string) (User, map[string]interface{}) {
	user, err := service.repo.InsertPassword(email, password)
	fmt.Println("User id:", user.ID)
	var response map[string]interface{}
	if err != "" {
		user.ID = 0
		response = apihelpers.Message(http.StatusNotFound, "User not found")
	}
	return user, response
}

func (service loginService) SignUp(user User) map[string]interface{} {
	result := service.repo.IsEmailExisting(user.Contact.Email)
	var response map[string]interface{}
	if result {
		return apihelpers.Message(http.StatusUnauthorized, "Email already exists")
	} else {
		user, err := service.repo.CreateUser(user)
		fmt.Println("User id:", user.ID)
		if err != "" {
			response = apihelpers.Message(0, err)
		} else {
			response = apihelpers.Message(http.StatusCreated, "User created successfully")
		}
		// Clear password to avoid sending it back to frontend
		user.Contact.Password = ""
		response["data"] = user
		return response
	}
}
