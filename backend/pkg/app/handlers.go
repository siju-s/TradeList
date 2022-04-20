package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"tradelist/pkg/api"
	"tradelist/pkg/apihelpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/mapstructure"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

var jwtKey = []byte("secret_key")

type Credentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Claims struct {
	Username           string `json:"username"`
	jwt.StandardClaims        //go get github.com/golang-jwt/jwt

}

func (server *Server) Signup(writer http.ResponseWriter, request *http.Request) {
	var user api.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := server.loginService.SignUp(user)
	apihelpers.Respond(writer, response)
}

func (server *Server) Login(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user, response := server.loginService.FetchUserInfo(credentials.Email)

	if user.ID == 0 {
		apihelpers.Respond(writer, response)
		return
	}

	expectedPassword := user.Contact.Password

	if err := bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(credentials.Password)); err != nil {
		response = apihelpers.Message(http.StatusUnauthorized, "Incorrect Password")
		apihelpers.Respond(writer, response)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		response = apihelpers.Message(http.StatusInternalServerError, "Server error")
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	user.Contact.Password = ""
	response["data"] = user
	apihelpers.Respond(writer, response)
}

func (server *Server) Home(writer http.ResponseWriter, request *http.Request) {
	c, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	writer.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

}

func (server *Server) Refresh(writer http.ResponseWriter, request *http.Request) {
	c, err := request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	//commenting out for testing purposes
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `refresh_token` cookie
	http.SetCookie(writer, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func (server *Server) Logout(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(writer, &cookie)

	writer.Write([]byte("Old cookie deleted. Logged out!\n"))

}

func (server *Server) ForgotPassword(writer http.ResponseWriter, request *http.Request) {

	var data map[string]string // string as a key and string as a value
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	email := data["Email"]
	token := GenerateRandomString(12)

	user, response := server.loginService.InsertToken(data["Email"], token)
	user.Token = token

	if response != nil {
		apihelpers.Respond(writer, response)
		return
	}

	from := "admin@example.com"

	to := []string{
		email,
	}
	//url := "http://localhost:4200/reset/" + token
	message := []byte("Token to reset password: " + token)

	smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

	response = apihelpers.Message(http.StatusOK, "Token sent! Check your mail")
	apihelpers.Respond(writer, response)

}

func GenerateRandomString(n int) string {
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	ret := make([]rune, n)

	for i := range ret {
		ret[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(ret)
}

func (server *Server) ResetPassword(writer http.ResponseWriter, request *http.Request) {

	var data map[string]string // string as a key and string as a value
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, response := server.loginService.VerifyToken(data["Token"])

	if response != nil {
		apihelpers.Respond(writer, response)
		return
	}

	if data["Password"] != data["Password_confirm"] {

		writer.Write([]byte("Passwords do not match!\n"))
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user, response = server.loginService.InsertPassword(data["Email"], string(password))

	if response != nil {
		apihelpers.Respond(writer, response)
		return
	}
	user.Contact.Password = string(password)

	response = apihelpers.Message(http.StatusOK, "Password Reset Successfully!")
	apihelpers.Respond(writer, response)

}

//TODO 1. Read post data correctly DONE
// 2. Upload image to AWS DONE
// 3. Save image url in DB DONE
// 4. Verify any user can upload images
func (server *Server) CreatePost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	err := request.ParseMultipartForm(32 << 20)
	if err != nil {
		return
	}
	var post api.Post
	var result map[string]interface{}
	categoryId, err := strconv.Atoi(mux.Vars(request)["id"])
	var jobPost api.JobPost

	for key, value := range request.Form {
		fmt.Printf("%s = %s\n\n", key, value)
		json.Unmarshal([]byte(value[0]), &result)
		mapstructure.Decode(result["Post"], &post)

		if categoryId == 1 {
			mapstructure.Decode(result, &jobPost)
			post = jobPost.Post
		} else {
			sendErr(writer, http.StatusBadRequest, "Invalid categoryid:"+strconv.Itoa(categoryId))
			return
		}
	}
	images := uploadImages(writer, request, post.SellerId)
	jobPost.Post.Image = images
	post.Image = images

	fmt.Println(jobPost.Post.Image)

	if err != nil {
		sendErr(writer, http.StatusBadRequest, err.Error())
		return
	}
	var response map[string]interface{}
	if categoryId == 1 {
		response = server.jobService.CreateJobPost(jobPost)
	} else {
		response = server.PostService.Create(post)
	}
	apihelpers.Respond(writer, response)
}

func uploadImages(writer http.ResponseWriter, request *http.Request, sellerid int) []api.Images {
	filelist := UploadHandler(writer, request, sellerid)

	var images []api.Images
	for _, path := range filelist {
		var image api.Images
		image.Url = path
		image.SellerId = sellerid
		images = append(images, image)
	}
	return images
}

func (server *Server) GetAllPosts(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := server.PostService.GetAllPosts()
	apihelpers.Respond(writer, response)
}

func (server *Server) GetPostsByUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	user_id := mux.Vars(request)["id"]
	fmt.Println("user_id:", user_id)
	response := server.PostService.GetPostsByUser(user_id)
	apihelpers.Respond(writer, response)
}

func (server *Server) GetAllCategories(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := server.PostService.GetAllCategories()
	apihelpers.Respond(writer, response)
}

func (server *Server) GetLocations(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	response := server.PostService.GetLocations()
	apihelpers.Respond(writer, response)
}

func (server *Server) GetSubcategories(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(request)["id"]
	fmt.Println("CategoryId:", categoryId)
	response := server.PostService.GetSubcategories(categoryId)
	apihelpers.Respond(writer, response)
}

func (server *Server) GetPostById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	postId := mux.Vars(request)["id"]
	fmt.Println("PostId:", postId)
	response := server.PostService.GetPostById(postId)
	apihelpers.Respond(writer, response)
}

func (server *Server) UpdatePost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	postId := mux.Vars(request)["id"]
	userId := mux.Vars(request)["userid"]

	fmt.Println("PostId:", postId+" user id: ", userId)

	var post api.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		sendErr(writer, http.StatusBadRequest, err.Error())
		return
	}
	response := server.PostService.UpdatePost(post, postId, userId)
	apihelpers.Respond(writer, response)
}

func (server *Server) DeletePost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	postId := mux.Vars(request)["id"]
	userId := mux.Vars(request)["userid"]

	fmt.Println("PostId:", postId+" user id: ", userId)
	response := server.PostService.DeletePost(postId, userId)
	apihelpers.Respond(writer, response)
}

func (server *Server) GetPostByCategoryId(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(request)["id"]
	fmt.Println("CategoryId:", categoryId)
	response := server.jobService.GetPostByCategoryId(categoryId)
	apihelpers.Respond(writer, response)
}

func (server *Server) GetPostBySubcategoryId(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	subcategoryId := mux.Vars(request)["id"]
	fmt.Println("subcategoryId:", subcategoryId)
	response := server.jobService.GetPostBySubcategoryId(subcategoryId)
	apihelpers.Respond(writer, response)
}

func UploadHandler(w http.ResponseWriter, r *http.Request, userid int) []string {
	files := r.MultipartForm.File["files"]

	var S3AccessId = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	var S3Secret = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	var S3Region = GetEnvWithKey("AWS_REGION")

	newSession, err := session.NewSession(&aws.Config{
		Region: aws.String(S3Region),
		Credentials: credentials.NewStaticCredentials(
			S3AccessId,
			S3Secret,
			"")})
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Could not upload files")
		return nil
	}

	var filenames []string

	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Could not get uploaded file")
			return nil
		}

		fileName, err := UploadFileToS3(newSession, file, header, userid)
		if err != nil {
			fmt.Fprintf(w, "Could not upload file")
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Println(w, "Image uploaded successfully: %v", fileName)
		}
		file.Close()
		filenames = append(filenames, fileName)
	}
	return filenames
}

func UploadFileToS3(s *session.Session, file multipart.File, header *multipart.FileHeader, userid int) (string, error) {
	size := header.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	var S3Bucket = GetEnvWithKey("AWS_BUCKET")

	// create a unique file name for the file
	tempFileName := "pictures/" + strconv.Itoa(userid) + "/" + bson.NewObjectId().Hex() + filepath.Ext(header.Filename)

	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3Bucket),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return tempFileName, err
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
