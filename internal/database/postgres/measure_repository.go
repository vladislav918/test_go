package repository

import (
	"test-project/internal/models"

	"gorm.io/gorm"
)

type MeasureRepository struct {
	DB *gorm.DB
}

func NewMeasureRepository(db *gorm.DB) *MeasureRepository {
	return &MeasureRepository{DB: db}
}

func (r *MeasureRepository) FindAll() ([]models.Measure, error) {
	var measures []models.Measure
	if result := r.DB.Find(&measures); result.Error != nil {
		return nil, result.Error
	}
	return measures, nil
}

func (r *MeasureRepository) FindByID(id int) (*models.Measure, error) {
	var measure models.Measure
	if result := r.DB.First(&measure, id); result.Error != nil {
		return nil, result.Error
	}
	return &measure, nil
}

func (r *MeasureRepository) Create(measure *models.Measure) error {
	return r.DB.Create(measure).Error
}

func (r *MeasureRepository) Update(measure *models.Measure) error {
	return r.DB.Save(measure).Error
}

func (r *MeasureRepository) Delete(id int) error {
	return r.DB.Delete(&models.Measure{}, id).Error
}
