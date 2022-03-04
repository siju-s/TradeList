package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tradelist/pkg/api"
	"tradelist/pkg/app"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
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
