package domain

import (
	"github.com/nqmt/go-service/handler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		product := &Product{
			ID:          "1",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}
		product.Update("2", "2", "2", 2, 2)

		expect := &Product{
			ID:          "1",
			Name:        "2",
			Description: "2",
			Image:       "2",
			Price:       2,
			Quantity:    2,
		}
		require.Equal(t, expect, product)
	})

	t.Run("nil Product", func(t *testing.T) {
		product := &Product{}
		product.Update("2", "2", "2", 2, 2)

		expect := &Product{
			ID:          "",
			Name:        "2",
			Description: "2",
			Image:       "2",
			Price:       2,
			Quantity:    2,
		}
		require.Equal(t, expect, product)
	})
}

func TestProduct_ToOutput(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		product := &Product{
			ID:          "1",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}

		got := product.ToOutput()
		expect := &handler.OutputCreateProduct{
			ID:          "1",
			Name:        "1",
			Description: "1",
			Image:       "1",
			Price:       1,
			Quantity:    1,
		}

		require.Equal(t, expect, got)
	})

	t.Run("nil Product", func(t *testing.T) {
		product := &Product{}

		got := product.ToOutput()
		expect := &handler.OutputCreateProduct{
			ID:          "",
			Name:        "",
			Description: "",
			Image:       "",
			Price:       0,
			Quantity:    0,
		}

		require.Equal(t, expect, got)
	})
}
