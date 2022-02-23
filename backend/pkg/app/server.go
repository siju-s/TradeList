package app

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"tradelist/pkg/api"
)

type Server struct {
	router      *mux.Router
	postService api.PostService
	jobService  api.JobService
}

func CreateServer(router *mux.Router, postService api.PostService, jobService api.JobService) *Server {
	return &Server{
		router:      router,
		postService: postService,
		jobService:  jobService}
}

func (server *Server) Run() {
	server.Routes()
	handler := cors.AllowAll().Handler(server.router)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
