package app

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"tradelist/pkg/api"
)

type Server struct {
	Router       *mux.Router
	PostService  api.PostService
	jobService   api.JobService
	loginService api.LoginService
}

func CreateServer(router *mux.Router, postService api.PostService, jobService api.JobService, loginService api.LoginService) *Server {
	return &Server{
		Router:       router,
		PostService:  postService,
		jobService:   jobService,
		loginService: loginService,
	}
}

func (server *Server) Run() {
	server.Routes()
	handler := cors.AllowAll().Handler(server.Router)
	log.Fatal(http.ListenAndServe(":8081", handler))
}

func (server *Server) RunTest() {
	server.Routes()
	cors.AllowAll().Handler(server.Router)
}
