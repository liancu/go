package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" binding:"required,min=2,max=100,lowercase"`
	Price     uint           `json:"price" binding:"required"`
}
