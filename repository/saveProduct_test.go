package repository

import (
	"github.com/nqmt/go-service/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_SaveProduct(t *testing.T) {
	ts := NewTestSuite()

	input := &domain.Product{
		ID:          "1",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}

	product, err := ts.repo.SaveProduct(input)

	expect := &domain.Product{
		ID:          "1",
		Name:        "1",
		Description: "1",
		Image:       "1",
		Price:       1,
		Quantity:    1,
	}

	require.NoError(t, err)
	require.Equal(t, expect, product)

	ts.TearDown()
}
