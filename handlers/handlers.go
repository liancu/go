package handlers

import (
	"fmt"
	"github.com/adore-me/hello/model"
	"github.com/adore-me/hello/repository"
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
	var request productRequest
	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	repo := ctx.MustGet("repo").(*repository.SqlLite)

	p, err := repo.Create(&model.Product{
		Name:  request.Name,
		Price: request.Price,
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}

func GetProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	// .(*repository.SqlLite) - cast!
	repo := ctx.MustGet("repo").(*repository.SqlLite)

	p, err := repo.GetOne(productId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}
func GetAllProducts(ctx *gin.Context) {
	// .(*repository.SqlLite) - cast!
	repo := ctx.MustGet("repo").(*repository.SqlLite)

	products, err := repo.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": products})
}

func UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	var request productRequest
	err := ctx.BindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	repo := ctx.MustGet("repo").(*repository.SqlLite)

	//todo send an array of column and values
	p, err := repo.Update(productId, model.Product{
		Name:  request.Name,
		Price: request.Price,
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": p})
}

func DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")

	repo := ctx.MustGet("repo").(*repository.SqlLite)

	err := repo.Delete(productId)

	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
