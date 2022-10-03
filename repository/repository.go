package repository

import (
	"github.com/adore-me/hello/model"
)

type Product interface {
	Create(*model.Product) (*model.Product, error)
	Get() (*[]model.Product, error)
	GetOne(productId string) (*model.Product, error)
	Update(productId string, productData model.Product) (*model.Product, error)
	Delete(productId string) error
}
