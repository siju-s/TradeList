package app

func (server *Server) Routes() {
	router := server.router
	router.HandleFunc("/post", server.CreatePost).Methods("POST", "OPTIONS")
	router.HandleFunc("/post", server.GetAllPosts).Methods("GET", "OPTIONS")
	router.HandleFunc("/post/{id}", server.GetPostById).Methods("GET")
	router.HandleFunc("/post/{id}", server.UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", server.DeletePost).Methods("DELETE")

	router.HandleFunc("/categories", server.GetAllCategories).Methods("GET")
	router.HandleFunc("/subcategories/{id}", server.GetSubcategories).Methods("GET")

	router.HandleFunc("/post/category/{id}", server.CreateJobPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/post/category/{id}", server.GetPostByCategoryId).Methods("GET", "OPTIONS")
}
