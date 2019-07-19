package domain

import "github.com/nqmt/go-service/handler"

type Product struct {
	ID          string `bson:"_id"`
	Name        string
	Description string
	Image       string
	Price       float32
	Quantity    int
}

func (p *Product) Update(name, description, image string, price float32, quantity int) {
	p.Name = name
	p.Description = description
	p.Image = image
	p.Price = price
	p.Quantity = quantity
}

func (p *Product) ToOutput() *handler.OutputCreateProduct {
	return &handler.OutputCreateProduct{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}
}
