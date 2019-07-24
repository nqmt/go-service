package product

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//go:generate mockery -name=IProductRepo
type IProductRepo interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
	GetProduct(id string) *Product
	GetProductList() ([]*Product,error)
}

type ProductRepo struct {
	product *mongo.Collection
}

func NewProductRepo(mongo *mongo.Client, database, collection string) *ProductRepo {
	return &ProductRepo{
		product: mongo.Database(database).Collection(collection),
	}
}

func (p *ProductRepo) CreateProduct(product *Product) error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := p.product.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepo) UpdateProduct(product *Product) error {
	return nil
}

func (p *ProductRepo) DeleteProduct(id string) error {
	return nil
}

func (p *ProductRepo) GetProduct(id string) *Product {
	return nil
}

func (p *ProductRepo) GetProductList() ([]*Product,error) {
	return nil, nil
}