package app

import (
	"fmt"
	"gorm.io/gorm"
	"tradelist/pkg/api"
)

func (server *Server) Migrate(db *gorm.DB) {
	err := db.AutoMigrate(api.Contact{}, api.Category{}, api.Subcategory{}, api.User{}, api.Seller{},
		api.Images{}, api.Post{},
		api.PayType{}, api.JobType{}, api.Places{}, api.LocationType{}, api.Job{})
	if err != nil {
		return
	}
	result := db.Exec("PRAGMA foreign_keys = ON", nil)
	if result.Error != nil {
		print(result.Error)
	}
	if isCategoryNotExisting(db) {
		createDefaultValues(db)
	}
}

func isCategoryNotExisting(db *gorm.DB) bool {
	var categories []api.Category

	result := db.Find(&categories)
	fmt.Println("Categories rows:", result.RowsAffected)
	if result.RowsAffected > 0 {
		return false
	} else {
		return true
	}
}

func createDefaultValues(db *gorm.DB) {
	var categories = []api.Category{{Name: "Jobs"}, {Name: "Property"}, {Name: "For Sale"}}
	result := db.Create(&categories)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var subcategories = []api.Subcategory{
		{CategoryId: 1, Name: "Accounting"},
		{CategoryId: 1, Name: "HR"},
		{CategoryId: 1, Name: "Legal"},
		{CategoryId: 1, Name: "Customer Service"},
		{CategoryId: 1, Name: "Healthcare"},
		{CategoryId: 1, Name: "Hospitality"},
		{CategoryId: 1, Name: "Housekeeping"},
		{CategoryId: 1, Name: "Software"},
		{CategoryId: 1, Name: "Accounting"},

		{CategoryId: 2, Name: "For Sale"},
		{CategoryId: 2, Name: "To Rent"},
		{CategoryId: 2, Name: "To Share"},
		{CategoryId: 2, Name: "Sublet"},
		{CategoryId: 2, Name: "Storage"},

		{CategoryId: 3, Name: "Appliances"},
		{CategoryId: 3, Name: "Audio equipment"},
		{CategoryId: 3, Name: "Books"},
		{CategoryId: 3, Name: "Clothes"},
		{CategoryId: 3, Name: "Computers"},
		{CategoryId: 3, Name: "Furniture"},
		{CategoryId: 3, Name: "Gym equipment"},
		{CategoryId: 3, Name: "Sports equipment"},
	}
	db.Create(&subcategories)

	var seller = api.Seller{
		Contact: api.Contact{FirstName: "Test", LastName: "Seller"},
	}
	db.Create(&seller)
	createDefaultJobValues(db)
}

func createDefaultJobValues(db *gorm.DB) {
	var paytype = []api.PayType{
		{"monthly"},
		{"yearly"},
		{"hourly"},
	}
	var locationtype = []api.LocationType{
		{"remote"},
		{"onsite"},
		{"hybrid"},
	}
	var jobtype = []api.JobType{
		{"fulltime"},
		{"parttime"},
		{"internship"},
	}
	var place = []api.Places{
		{"Gainesville"},
		{"California"},
		{"Chicago"},
		{"NewYork"},
	}
	result := db.Create(&paytype)
	result = db.Create(&locationtype)
	result = db.Create(&jobtype)
	result = db.Create(&place)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
