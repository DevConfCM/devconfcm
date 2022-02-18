package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func GetUser(ID int) *User {
	var user User
	db.First(&user, ID)
	return &user
}

func CreateUser(u *User) *User {
	db.Create(u)
	return u
}
