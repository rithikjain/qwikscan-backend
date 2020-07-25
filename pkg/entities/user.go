package entities

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ImageUrl    string `json:"image_url"`
	PhoneNumber string `json:"phone_number"`
}
