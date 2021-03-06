package repository

import (
	"github.com/teten-nugraha/go-dev-forum/config"
	"github.com/teten-nugraha/go-dev-forum/models"
)

func FindUserById(id uint) models.User {
	var user models.User

	config.DB.Where("id = ?", id).Find(&user)

	return user
}

func FindUserByEmail(email string) models.User {

	var user models.User

	config.DB.Where("email = ?", email).Find(&user)

	return user
}

func CreateUser(user models.User) models.User {
	config.DB.Create(&user)
	return user
}
