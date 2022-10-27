package repository

import (
	"github.com/adore-me/hello/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlLite struct {
	db *gorm.DB
}

// Create SqlLite este receiver (asta inseamna ca metodele tin de SQLite repo)
func (r SqlLite) Create(product *model.Product) (*model.Product, error) {
	err := r.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r SqlLite) Get() (*[]model.Product, error) {
	var p []model.Product
	//todo inseamna ca Find-ul pune in variabla p rezultatul metodei? da
	err := r.db.Find(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r SqlLite) GetOne(productId string) (*model.Product, error) {
	var p model.Product
	err := r.db.First(&p, productId).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r SqlLite) Update(productId string, productData model.Product) (*model.Product, error) {
	var p model.Product
	err := r.db.Model(&p).Where("id = ?", productId).Updates(productData).Error

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r SqlLite) Delete(productId string) error {
	var p model.Product
	return r.db.Delete(&p, productId).Error
}

func NewSqlLite(dsn string) (*SqlLite, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SqlLite{db: db}, nil
}
