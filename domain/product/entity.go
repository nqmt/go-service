package product

type Product struct {
	ID          string  `json:"id" bson:"_id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" binding:"gte=0"`
	Quantity    int     `json:"quantity" binding:"gte=0"`
}

func (p *Product) Update(name, description, image string, price float64, quantity int) {
	p.Name = name
	p.Description = description
	p.Image = image
	p.Price = price
	p.Quantity = quantity
}
