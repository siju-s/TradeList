package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"tradelist/pkg/api"
	"tradelist/pkg/app"
)

func main() {
	LoadEnv()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error())
	}

	repo := api.CreateRepo(db)

	postService := api.CreatePostService(repo)
	jobService := api.CreateJobService(repo)
	loginService := api.CreateLoginService(repo)
	server := app.CreateServer(mux.NewRouter(), postService, jobService, loginService)
	server.Migrate(db)

	server.Run()
}

// Required for loading AWS secrets. Make sure the .env file is present under backend\cmd folder
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}
}
