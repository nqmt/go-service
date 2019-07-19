package repository

import (
	"context"
	"github.com/nqmt/go-service/domain"
)

func (r *Repository) SaveProduct(product *domain.Product) (*domain.Product, error) {
	_, err := r.product.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
