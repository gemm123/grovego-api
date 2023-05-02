package service

import (
	"errors"
	"gemm123/grovego-api/helper"
	"gemm123/grovego-api/jwt"
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
	CheckAccount(input models.Login) error
	GenerateToken(input models.Login) (string, error)
	GetUser(userID string) (models.UserResponse, error)
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

func (s *serviceUser) CheckAccount(input models.Login) error {
	user, err := s.repositoryUser.FindUserByEmail(input.Email)
	if err != nil {
		return errors.New("email or password not registered")
	}

	ok := helper.CheckPasswordHash(input.Password, user.Password)
	if !ok {
		return errors.New("email or password not registered")
	}

	return err
}

func (s *serviceUser) GenerateToken(input models.Login) (string, error) {
	user, err := s.repositoryUser.FindUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.ID, user.Email, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *serviceUser) GetUser(userID string) (models.UserResponse, error) {
	user, err := s.repositoryUser.FindUserByID(userID)
	if err != nil {
		return models.UserResponse{}, err
	}

	userResponse := models.UserResponse{
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Phone:     user.Phone,
		ImagePath: user.ImagePath,
	}

	return userResponse, nil
}
