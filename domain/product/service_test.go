package product_test

import (
	"github.com/nqmt/go-service/domain/product"
	"github.com/nqmt/go-service/domain/product/mocks"
	"github.com/nqmt/go-service/util/xidgen"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestSuite struct {
	mockRepo *mocks.IProductRepo
	Service  *product.Service
}

func (ts *TestSuite) Mock(mockType string) {
	switch mockType {
	case "SaveProduct_Success":
		ts.mockRepo.On("CreateProduct", &product.Product{
			ID:          "xxx",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}).Return(nil)
	}
}

func NewTestSuite() *TestSuite {
	m := &mocks.IProductRepo{}

	return &TestSuite{
		mockRepo: m,
		Service:  product.NewService(m),
	}
}

func TestService_CreateProduct1(t *testing.T) {
	m := &mocks.IProductRepo{}
	service := product.NewService(m)

	pid := xidgen.New()
	pid.Freeze("xxx")

	m.On("CreateProduct", &product.Product{
		ID:          "xxx",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}).Return(nil)

	got, err := service.CreateProduct(&product.Product{
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	})

	expect := &product.Product{
		ID:          "xxx",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}

	require.Equal(t, expect, got)
	require.NoError(t, err)
}

// Testing: the Arrange, Act and Assert Pattern
func TestService_CreateProduct2(t *testing.T) {
	ts := NewTestSuite()
	ts.Mock("SaveProduct_Success")

	got, err := ts.Service.CreateProduct(&product.Product{
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	})

	expect := &product.Product{
		ID:          "xxx",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}

	require.Equal(t, expect, got)
	require.NoError(t, err)

}
