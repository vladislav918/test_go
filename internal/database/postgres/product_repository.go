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
	var products []models.Product
	if result := r.DB.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	var product models.Product
	if result := r.DB.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(id int) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
