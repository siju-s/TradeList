package app

import (
	"fmt"
	"gorm.io/gorm"
	"tradelist/pkg/api"
	"tradelist/pkg/data"
)

func (server *Server) Migrate(db *gorm.DB) {
	err := db.AutoMigrate(api.Contact{}, api.Category{}, api.Subcategory{}, api.User{}, api.Seller{},
		api.Images{}, api.Post{},
		api.PayType{}, api.JobType{}, api.Places{}, api.LocationType{}, api.Job{})
	db.Migrator().RenameColumn(api.User{}, "id", "user_id")
	db.Migrator().RenameColumn(api.Job{}, "id", "job_id")

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
	var categories = []api.Category{{Name: "Jobs"}, {Name: "Property"}, {Name: "For Sale"}, {Name: "Services"}, {Name: "Community"}}
	result := db.Create(&categories)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	var subcategories = data.GetSubcategories()
	db.Create(&subcategories)
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
