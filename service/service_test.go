package service

import (
	"github.com/nqmt/go-service/domain"
	"github.com/nqmt/go-service/repository/mocks"
)

type TestSuite struct {
	mockRepo *mocks.IRepository
	Service *Service
}

func (ts *TestSuite) Mock(mockType string) {
	switch mockType {
	case "SaveProduct_Success":
		ts.mockRepo.On("SaveProduct",  &domain.Product{
			ID:          "xxx",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}).Return(&domain.Product{
			ID:          "xxx",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}, nil)
	}
}

func NewTestSuite() *TestSuite {
	m := &mocks.IRepository{}
	return &TestSuite{
		mockRepo: m,
		Service: New(m),
	}
}
