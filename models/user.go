package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err

	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]User) error {
	err := db.Find(User).Error
	if err != nil {
		return err

	}
	return nil
}

func GetUser(db *gorm.DB, User *User, id int) error {
	err := db.Debug().Where("id = ?", id).First(User).Error
	if err != nil {
		return err

	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) error {
	err := db.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(db *gorm.DB, User *User, id int) error {
	db.Where("id = ?", id).Delete(User)
	return nil
}
