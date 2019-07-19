package service

import (
	"github.com/nqmt/go-service/domain"
	"github.com/nqmt/go-service/handler"
	"github.com/nqmt/go-service/repository/mocks"
	"github.com/nqmt/go-service/util/id"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_CreateProduct1(t *testing.T) {
	m := &mocks.IRepository{}
	service := New(m)

	pid := id.New()
	pid.Freeze("xxx")

	m.On("SaveProduct", &domain.Product{
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

	got, err := service.CreateProduct(&handler.InputCreateProduct{
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	})

	expect := &handler.OutputCreateProduct{
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

	got, err := ts.Service.CreateProduct(&handler.InputCreateProduct{
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	})

	expect := &handler.OutputCreateProduct{
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