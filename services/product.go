package services

import (
	"github.com/adore-me/hello/model"
	"github.com/adore-me/hello/repository"
	"github.com/sirupsen/logrus"
)

type Product struct {
	repo   repository.Product
	logger *logrus.Logger
}

func NewProductService(repo repository.Product, logger *logrus.Logger) *Product {
	return &Product{
		repo:   repo,
		logger: logger,
	}
}

func (p *Product) CreateProduct(prod *model.Product) (*model.Product, error) {
	p.logger.Info("se creaza un produs")

	return p.repo.Create(prod)
}

func (p *Product) GetProduct(productId string) (*model.Product, error) {
	p.logger.Info("get product")

	return p.repo.GetOne(productId)
}

func (p *Product) GetProducts() (*[]model.Product, error) {
	p.logger.Info("get products")

	return p.repo.Get()
}

func (p *Product) UpdateProduct(productId string, productData model.Product) (*model.Product, error) {
	p.logger.Info("update product")

	return p.repo.Update(productId, productData)
}

func (p *Product) DeleteProduct(productId string) error {
	p.logger.Info("delete product")

	return p.repo.Delete(productId)
}
