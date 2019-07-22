package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nqmt/go-service/product"
	"github.com/nqmt/go-service/soap"
	"github.com/nqmt/go-service/util/con"
)

func main() {
	g := gin.New()

	// root middleware
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(gzip.Gzip(gzip.DefaultCompression))

	// connect db
	mongo := con.ConnectMongo("mongodb://test:test@localhost:27017")
	oracle := con.ConnectOracle()

	// repository
	productRepo := product.NewProductRepo(mongo, "e-commerce", "product")
	soapRepo := soap.NewSoapRepo(oracle)

	// new service
	sproduct := product.NewService(productRepo)
	ssoap := soap.NewService(soapRepo)

	// handler
	productHandler := product.NewHandler(g, sproduct)
	productHandler.RegisterRouter()

	soapHandler := soap.NewHandler(g, ssoap)
	soapHandler.RegisterRouter()

	err := g.Run(":6001")
	if err != nil {
		panic(err)
	}
}