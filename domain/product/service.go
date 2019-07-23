package product

import (
	"github.com/nqmt/go-service/util/errorx"
	"github.com/nqmt/go-service/util/xidgen"
)

type IService interface {
	CreateProduct(product *Product) (*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
	GetProduct(id string) (*Product, error)
	GetProductList() ([]*Product, error)
}

type Service struct {
	productRepo IProductRepo
}

func NewService(product IProductRepo) *Service {
	return &Service{
		productRepo: product,
	}
}

func (s *Service) CreateProduct(input *Product) (*Product, error) {
	xid := xidgen.New()

	product := &Product{
		ID:          xid.Gen(),
		Name:        input.Name,
		Description: input.Description,
		Image:       input.Image,
		Price:       input.Price,
		Quantity:    input.Quantity,
	}

	err := s.productRepo.CreateProduct(product)
	if err != nil {
		// ErrUnableCreateProduct(err)
		return nil, errorx.InternalServerError(err.Error(), "UNABLE_CREATE_PRODUCT", "unable create product")
	}

	return product, nil
}

func (s *Service) UpdateProduct(input *Product) error {
	return nil
}

func (s *Service) DeleteProduct(id string) error {
	return nil
}

func (s *Service) GetProduct(id string) (*Product, error) {
	return nil, nil
}

func (s *Service) GetProductList() ([]*Product, error) {
	return nil, nil
}
