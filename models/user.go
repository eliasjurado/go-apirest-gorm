package models

import (
	"apirest-gorm/database"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrateUser() {
	db, err := database.Connect()
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	db.AutoMigrate(User{})
}
