package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}

func (a *App) start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
