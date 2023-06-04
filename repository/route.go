package repository

import (
	"gemm123/grovego-api/models"

	"gorm.io/gorm"
)

type repositoryRoute struct {
	DB *gorm.DB
}

type RepositoryRoute interface {
	CreateRouteUser(route models.Route) error
}

func NewRepositoryRoute(DB *gorm.DB) *repositoryRoute {
	return &repositoryRoute{DB: DB}
}

func (r *repositoryRoute) CreateRouteUser(route models.Route) error {
	err := r.DB.Create(&route).Error
	return err
}
