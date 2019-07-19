package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestSuite struct {
	mongo *mongo.Client
	repo *Repository
}


func NewTestSuite() *TestSuite {
	mongo := ConnectMongoDB("mongodb://test:test@localhost:27017")

	return &TestSuite{
		mongo: mongo,
		repo: New(mongo),
	}
}

func (ts *TestSuite) TearDown() {
	err := ts.mongo.Database("e-commerce").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}