package main

import (
	"log"
	"os"
	"time"
	"tradelist/pkg/api"
	"tradelist/pkg/app"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//LoadEnv()
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
	jobService := api.CreateJobService(db)
	server := app.CreateServer(mux.NewRouter(), postService, jobService)
	server.Migrate(db)

	server.Run()
}

/*
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
*/
