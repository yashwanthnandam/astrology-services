package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type DashaRepository struct {
	db *gorm.DB
}

func NewDashaRepository(db *gorm.DB) repositories.DashaRepository {
	return &DashaRepository{db}
}

func (r *DashaRepository) GetDashaByID(id uint) (*models.Dasha, error) {
	var dasha models.Dasha
	result := r.db.First(&dasha, id)
	return &dasha, result.Error
}

func (r *DashaRepository) CreateDasha(dasha *models.Dasha) error {
	return r.db.Create(dasha).Error
}

func (r *DashaRepository) UpdateDasha(dasha *models.Dasha) error {
	return r.db.Save(dasha).Error
}

func (r *DashaRepository) DeleteDasha(id uint) error {
	return r.db.Delete(&models.Dasha{}, id).Error
}
