package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	app     *gin.Engine
	service *Service
}

func NewHandler(app *gin.Engine, service *Service) *Handler {
	return &Handler{
		app:     app,
		service: service,
	}
}

func (h *Handler) RegisterRouter() {
	h.app.POST("/product", h.CreateProduct())
}

func (h *Handler) CreateProduct() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := &Product{}

		err := c.BindJSON(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		output, err := h.service.CreateProduct(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, output)
	}
}