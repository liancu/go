package repository

import (
	"database/sql"
	"github.com/adore-me/hello/model"
)

type Product interface {
	Create(*model.Product) error
	Get(*model.Product) error
	GetOne(*model.Product, string) (*sql.Row, error)
	Update(product *model.Product, column string, value any) error
	Delete(*model.Product, int) error
}
