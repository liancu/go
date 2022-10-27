package handlers

import (
	"fmt"
	"github.com/adore-me/hello/model"
	"github.com/adore-me/hello/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func Ready(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func Hello(ctx *gin.Context) {
	//req: = helloRequest{} - aici ar si initializa, mai jos doar defineste
	var req helloRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "bine ai venit " + req.Name})

	fmt.Println(req)
}

func CreateProduct(ctx *gin.Context) {
	var prod model.Product
	err := ctx.BindJSON(&prod)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	service := ctx.MustGet("product-service").(*services.Product)

	p, err := service.CreateProduct(&prod)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}

func GetProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	service := ctx.MustGet("product-service").(*services.Product)

	p, err := service.GetProduct(productId)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}
func GetAllProducts(ctx *gin.Context) {
	service := ctx.MustGet("product-service").(*services.Product)

	p, err := service.GetProducts()

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}

func UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	var prod model.Product
	err := ctx.BindJSON(&prod)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	service := ctx.MustGet("product-service").(*services.Product)

	p, err := service.UpdateProduct(productId, prod)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}

func DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	service := ctx.MustGet("product-service").(*services.Product)

	err := service.DeleteProduct(productId)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
