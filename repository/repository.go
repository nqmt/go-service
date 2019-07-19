package repository

import (
	"context"
	"github.com/nqmt/go-service/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//go:generate mockery -name=IRepository
type IRepository interface {
	SaveProduct(product *domain.Product) (*domain.Product, error)
}

type Repository struct {
	product *mongo.Collection
}

func ConnectMongoDB(url string) *mongo.Client {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}

func New(mongo *mongo.Client) *Repository {
	return &Repository{
		product: mongo.Database("e-commerce").Collection("product"),
	}
}
