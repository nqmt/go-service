package handler

type InputCreateProduct struct {
	Name        string
	Description string
	Image       string
	Price       float32
	Quantity    int
}

func (i InputCreateProduct) Validate() {

}
