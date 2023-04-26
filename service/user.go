package service

import (
	"gemm123/grovego-api/helper"
	"gemm123/grovego-api/models"
	"gemm123/grovego-api/repository"
	"time"

	"github.com/google/uuid"
)

type serviceUser struct {
	repositoryUser repository.RepositoyUser
}

type ServiceUser interface {
	Register(input models.RegisterUser) error
}

func NewServiceUser(repositoryUser repository.RepositoyUser) *serviceUser {
	return &serviceUser{
		repositoryUser: repositoryUser,
	}
}

func (s *serviceUser) Register(input models.RegisterUser) error {
	hashedPass, err := helper.HashPassword(input.Password)
	if err != nil {
		return err
	}

	newUser := models.User{
		ID:        uuid.New(),
		Name:      input.Name,
		Username:  input.Username,
		Email:     input.Email,
		Phone:     input.Phone,
		Password:  hashedPass,
		ImagePath: "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repositoryUser.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
