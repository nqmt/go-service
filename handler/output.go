package handler

type OutputCreateProduct struct {
	ID          string
	Name        string
	Description string
	Image       string
	Price       float32
	Quantity    int
}