package main

import (
	"github.com/adore-me/hello/handlers"
	"github.com/adore-me/hello/repository"
	"github.com/adore-me/hello/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()

	repo, err := repository.NewSqlLite("test.db")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	productService := services.NewProductService(repo, log)

	r.Use(func(c *gin.Context) {
		c.Set("product-service", productService)
		c.Next()
	})

	r.GET("/healthz", handlers.Healthy)
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.GET("/products", handlers.GetAllProducts)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.Run(":80")

	//todo read about defer, cum se scriu if-urile
}
