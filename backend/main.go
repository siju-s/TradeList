package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	app := App{
		db:  db,
		mux: mux.NewRouter(),
	}
	app.start()
}
