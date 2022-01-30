package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	db  *gorm.DB
	mux *mux.Router
}

type Contact struct {
	Name  string
	Email string
	Phone string
}

type Seller struct {
	SellerId int     `gorm:"primaryKey"`
	Contact  Contact `gorm:"embedded"`
	Rating   int     `gorm:"default:0"`
}

type Category struct {
	CategoryId int `gorm:"primaryKey"`
	Name       string
}

type Subcategory struct {
	SubcategoryId int `gorm:"primaryKey"`
	CategoryId    int
	Category      Category
	Name          string
}

type Post struct {
	gorm.Model
	SellerId      int
	Seller        Seller
	CategoryId    int
	SubcategoryId int
	Category      Category
	Subcategory   Subcategory
	Title         string `gorm:"not null"`
	Description   string `gorm:"not null"`
	IsHidden      bool   `gorm:"default:false"`
	IsFlagged     bool   `gorm:"default:false"`
	IsDeleted     bool   `gorm:"default:false"`
	HasImage      bool   `gorm:"default:false"`
}

type User struct {
	ID       int  `gorm:"primaryKey;autoIncrement"`
	IsSeller bool `gorm:"default:false"`
	SellerId int
	Seller   Seller
	Contact  Contact `gorm:"embedded"`
}

//Struct for Products

type Product struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    string `json:"price"`
}

//Init products var as a slice Product struct
var products []Product

//Get all Products
func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}

//Get Single Product
func getProduct(w http.ResponseWriter, r *http.Request) {

}

//Create Product
func createProduct(w http.ResponseWriter, r *http.Request) {

}

//Update Product
func updateProduct(w http.ResponseWriter, r *http.Request) {

}

//Delete Product
func deleteProduct(w http.ResponseWriter, r *http.Request) {

}

func (app *App) start() {

	err := app.db.AutoMigrate(&Contact{}, &Category{}, &Subcategory{}, &User{}, &Seller{}, &Post{})
	if err != nil {
		return
	}

	result := app.db.Exec("PRAGMA foreign_keys = ON", nil)
	if result.Error != nil {
		print(result.Error)
		return
	}

	//Init Router
	r := mux.NewRouter()

	//Mock Data
	products = append(products, Product{ID: "1", Category: "job", Name: "SDE", Price: "100000"})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/product", getProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/product", createProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", deleteProduct).Methods("DELETE")

	app.mux.HandleFunc("/post", app.savePost).Methods("POST")
	app.mux.HandleFunc("/posts", app.getAllPosts).Methods("GET")
	app.mux.HandleFunc("/", app.getAllPosts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func (app *App) getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var all []Post
	err := app.db.Find(&all).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(all)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (app *App) savePost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		sendErr(writer, http.StatusBadRequest, err.Error())
		return
	}
	err = app.db.Save(&post).Error
	if err != nil {
		sendErr(writer, http.StatusInternalServerError, err.Error())
	} else {
		writer.WriteHeader(http.StatusCreated)
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
