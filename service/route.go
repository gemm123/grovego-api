package service

import (
	"gemm123/grovego-api/models"
	"gemm123/grovego-api/repository"
)

type serviceRoute struct {
	repositoryRoute repository.RepositoryRoute
}

type ServiceRoute interface {
	CreateRouteUser(route models.Route) error
}

func NewServiceRoute(repositoryRoute repository.RepositoryRoute) *serviceRoute {
	return &serviceRoute{
		repositoryRoute: repositoryRoute,
	}
}

func (s *serviceRoute) CreateRouteUser(route models.Route) error {
	err := s.repositoryRoute.CreateRouteUser(route)
	return err
}
