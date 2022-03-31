package api

import "gorm.io/gorm"

type Contact struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Phone     string
}

type Seller struct {
	Rating int `gorm:"default:0"`
}

type Category struct {
	CategoryId int `gorm:"primaryKey;autoIncrement"`
	Name       string
}

type Subcategory struct {
	SubcategoryId int `gorm:"primaryKey"`
	CategoryId    int
	Category      Category `json:"-"`
	Name          string
}

type Post struct {
	gorm.Model
	SellerId      int
	User          User `json:"-" gorm:"foreignKey:SellerId"`
	CategoryId    int
	SubcategoryId int
	Category      Category    `json:"-"`
	Subcategory   Subcategory `json:"-"`
	Title         string      `gorm:"not null"`
	Description   string      `gorm:"not null"`
	IsHidden      bool        `gorm:"default:false"`
	IsFlagged     bool        `gorm:"default:false"`
	IsDeleted     bool        `gorm:"default:false"`
	Image         []Images    `gorm:"foreignKey:PostId;references:ID"`
}

type User struct {
	ID       int     `gorm:"primaryKey;autoIncrement"`
	IsSeller bool    `gorm:"default:false"`
	Seller   Seller  `gorm:"embedded"`
	Contact  Contact `gorm:"embedded"`
}

type Images struct {
	ID       int `gorm:"primaryKey"`
	Url      string
	SellerId int
	PostId   int
}

type JobPost struct {
	Post Post
	Job  Job
}

type PayType struct {
	Name string `gorm:"primaryKey"`
}

type Job struct {
	ID            int  `gorm:"primaryKey"`
	PostId        uint `gorm:"not null"`
	Post          Post `json:"-"`
	SubcategoryId int
	Subcategory   Subcategory `json:"-"`
	Salary        float32
	Pay           string       `gorm:"default:yearly"`
	PayType       PayType      `json:"-" gorm:"foreignKey:Pay"`
	Type          string       `gorm:"default:fulltime"`
	JobType       JobType      `json:"-" gorm:"foreignKey:Type;references:Name"`
	Location      string       `gorm:"default:onsite"`
	LocationType  LocationType `json:"-" gorm:"foreignKey:Location;references:Name"`
	Place         string       `gorm:"default:Gainesville"`
	Places        Places       `json:"-" gorm:"foreignKey:Place;references:Name"`
}

type JobType struct {
	Name string `gorm:"primaryKey"`
}

type LocationType struct {
	Name string `gorm:"primaryKey"`
}

type Places struct {
	Name string `gorm:"primaryKey"`
}

type PasswordReset struct {
	ID    uint
	Email string
	Token string `gorm:"unique"`
}
