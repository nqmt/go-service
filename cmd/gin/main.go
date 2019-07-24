package main

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nqmt/go-service/config"
	"github.com/nqmt/go-service/domain/product"
	"github.com/nqmt/go-service/util/con"
)

func main() {
	g := gin.New()

	// root middleware
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(gzip.Gzip(gzip.DefaultCompression))

	conf := config.Setup()

	// connect db
	mongo := con.ConnectMongo(conf.Mongodb.Connection)
	//oracle := con.ConnectSql("goracle", conf.Oracle.Connection)

	// repository
	productRepo := product.NewProductRepo(mongo, "e-commerce", "product")
	//soapRepo := soap.NewSoapRepo(oracle)

	// new service
	sproduct := product.NewService(productRepo)
	//ssoap := soap.NewService(soapRepo)

	// handler
	productHandler := product.NewHandler(g, sproduct)
	productHandler.RegisterRouter()

	//soapHandler := soap.NewHandler(g, ssoap)
	//soapHandler.RegisterRouter()

	err := g.Run(":" +conf.Port)
	if err != nil {
		panic(err)
	}
}