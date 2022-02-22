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

func GetAllUsers() *[]User {
	var users []User
	db.Find(&users)
	return &users
}

func UpdateUser(ID int, u *User) (*User, error) {
	var user User
	result := db.First(&user, ID)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	user.Name = u.Name
	user.Email = u.Email
	user.Username = u.Username
	db.Save(&user)
	return &user, nil
}

func DeleteUser(ID int) error {
	result := db.Delete(&User{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}