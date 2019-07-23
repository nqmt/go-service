package product_test

import (
	"github.com/nqmt/go-service/domain/product"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		p := &product.Product{
			ID:          "1",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}
		p.Update("2", "2", "2", 2, 2)

		expect := &product.Product{
			ID:          "1",
			Name:        "2",
			Description: "2",
			Image:       "2",
			Price:       2,
			Quantity:    2,
		}
		require.Equal(t, expect, p)
	})

	t.Run("nil Product", func(t *testing.T) {
		p := &product.Product{}
		p.Update("2", "2", "2", 2, 2)

		expect := &product.Product{
			ID:          "",
			Name:        "2",
			Description: "2",
			Image:       "2",
			Price:       2,
			Quantity:    2,
		}
		require.Equal(t, expect, p)
	})
}
