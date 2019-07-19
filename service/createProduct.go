package service

import (
	"github.com/nqmt/go-service/domain"
	"github.com/nqmt/go-service/handler"
	"github.com/nqmt/go-service/util/id"
)

func (s *Service) CreateProduct(input *handler.InputCreateProduct) (*handler.OutputCreateProduct, error) {
	pid := id.New()

	product := &domain.Product{
		ID:          pid.Gen(),
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	product, err := s.repo.SaveProduct(product)
	if err != nil {
		return nil, err
	}

	return product.ToOutput(), nil
}
