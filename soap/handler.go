package soap

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
	h.app.GET("/soap", h.GetSoap())
}

func (h *Handler) GetSoap() func(c *gin.Context) {
	return func(c *gin.Context) {
		output, err := h.service.GetSoap()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, output)
	}
}