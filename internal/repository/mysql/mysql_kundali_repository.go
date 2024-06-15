package mysql

import (
	"astrology-services/internal/domain/models"

	"gorm.io/gorm"
)

type KundaliRepository struct {
	db *gorm.DB
}

func NewKundaliRepository(db *gorm.DB) *KundaliRepository {
	return &KundaliRepository{db}
}

func (r *KundaliRepository) GetKundaliByID(id uint) (*models.Kundali, error) {
	var kundali models.Kundali
	result := r.db.First(&kundali, id)
	return &kundali, result.Error
}
