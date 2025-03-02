package repository

import (
	"test-project/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var books []models.Product
	if result := r.DB.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *ProductRepository) Create(book *models.Product) error {
	return r.DB.Create(book).Error
}
