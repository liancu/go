package main

import (
	"github.com/adore-me/hello/handlers"
	"github.com/adore-me/hello/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//todo: creez un produs, il updatez, cer o lista si o afisez pe ecran cu fmt.println, si dupa il sterg
	//todo sa pornesc api cu gin cum am facut data trecuta: il creez prin api, il citesc din db si il return din api. + sters

	repo, err := repository.NewSqlLite("test.db")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	//functie anonima
	r.Use(func(c *gin.Context) {
		c.Set("repo", repo)
		c.Next()
	})

	r.GET("/healthz", handlers.Healthy)
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.GET("/products", handlers.GetAllProducts)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.Run(":80")
}
