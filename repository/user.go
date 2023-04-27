package repository

import (
	"gemm123/grovego-api/models"

	"gorm.io/gorm"
)

type repositoryUser struct {
	DB *gorm.DB
}

type RepositoyUser interface {
	CreateUser(user models.User) error
	FindUserByEmail(email string) (models.User, error)
}

func NewRepositoyUser(DB *gorm.DB) *repositoryUser {
	return &repositoryUser{DB: DB}
}

func (r *repositoryUser) CreateUser(user models.User) error {
	err := r.DB.Create(&user).Error
	return err
}

func (r *repositoryUser) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
