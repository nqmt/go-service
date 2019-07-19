package service

import (
	"github.com/nqmt/go-service/handler"
	"github.com/nqmt/go-service/repository"
)

type IService interface {
	CreateProduct(product *handler.InputCreateProduct) (*handler.OutputCreateProduct, error)
}

type Service struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) *Service {
	return &Service{
		repo: repo,
	}
}
