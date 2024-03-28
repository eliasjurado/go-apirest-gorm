package repository

import (
	"apirest-gorm/database"
	"apirest-gorm/models"
)

func GetAllUsers() (models.Users, error) {
	users := models.Users{}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	db.Find(&users)
	return users, nil
}

func GetOneUser(id int64) (*models.User, error) {
	user := &models.User{}

	db, err := database.Connect()
	if err != nil {
		return user, err
	}
	db.Find(&user)
	return user, nil
}

func SaveUser(user *models.User) (*models.User, error) {
	db, err := database.Connect()
	if err != nil {
		return user, err
	}
	db.Save(&user)
	return user, nil
}

func RemoveUser(user *models.User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	db.Delete(&user)
	return nil
}
