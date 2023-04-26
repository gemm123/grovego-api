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
}

func NewRepositoyUser(DB *gorm.DB) *repositoryUser {
	return &repositoryUser{DB: DB}
}

func (r *repositoryUser) CreateUser(user models.User) error {
	err := r.DB.Create(&user).Error
	return err
}
