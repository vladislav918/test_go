package repository

import (
	// "test-project/internal/models"
	"gorm.io/gorm"
)

type Product interface {

}

type Measure interface {
	
}


type Repository struct {
	Product
	Measure
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Product: NewProductRepository(db),
		Measure: NewMeasureRepository(db),
	}
}
