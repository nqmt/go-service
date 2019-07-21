package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nqmt/go-service/product"
	"github.com/nqmt/go-service/util/con"
)

func main() {
	g := gin.New()

	// root middleware
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(gzip.Gzip(gzip.DefaultCompression))

	// connect db
	mongo := con.Connect("mongodb://test:test@localhost:27017")

	// repository
	productRepo := product.NewProductRepo(mongo, "e-commerce", "product")

	// new service
	s := product.NewService(productRepo)

	// handler
	productHandler := product.NewHandler(g, s)
	productHandler.RegisterRouter()

	err := g.Run(":6001")
	if err != nil {
		panic(err)
	}
}