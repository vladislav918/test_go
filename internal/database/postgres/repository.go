package repository

import (
	"test-project/internal/models"

	"gorm.io/gorm"
)

type Product interface {
	FindAll() ([]models.Product, error)
	FindByID(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

type Measure interface {
	FindAll() ([]models.Measure, error)
	FindByID(id int) (*models.Measure, error)
	Create(measure *models.Measure) error
	Update(measure *models.Measure) error
	Delete(id int) error
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
