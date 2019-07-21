package product_test

import (
	"context"
	"github.com/nqmt/go-service/product"
	"github.com/nqmt/go-service/util/con"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type RepoTestSuite struct {
	mongo *mongo.Client
	repo *product.ProductRepo
}


func NewRepoTestSuite() *RepoTestSuite {
	mongo := con.Connect("mongodb://test:test@localhost:27017")

	return &RepoTestSuite{
		mongo: mongo,
		repo: product.NewProductRepo(mongo, "e-commerce", "product"),
	}
}

func (ts *RepoTestSuite) TearDown() {
	err := ts.mongo.Database("e-commerce").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}

func TestRepository_SaveProduct(t *testing.T) {
	ts := NewRepoTestSuite()

	input := &product.Product{
		ID:          "1",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}

	err := ts.repo.CreateProduct(input)

	require.NoError(t, err)

	ts.TearDown()
}
