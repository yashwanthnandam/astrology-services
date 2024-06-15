package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type AshtakvargaRepository struct {
	db *gorm.DB
}

func NewAshtakvargaRepository(db *gorm.DB) repositories.AshtakvargaRepository {
	return &AshtakvargaRepository{db}
}

func (r *AshtakvargaRepository) GetAshtakvargaByID(id uint) (*models.Ashtakvarga, error) {
	var ashtakvarga models.Ashtakvarga
	result := r.db.First(&ashtakvarga, id)
	return &ashtakvarga, result.Error
}

func (r *AshtakvargaRepository) CreateAshtakvarga(ashtakvarga *models.Ashtakvarga) error {
	return r.db.Create(ashtakvarga).Error
}

func (r *AshtakvargaRepository) UpdateAshtakvarga(ashtakvarga *models.Ashtakvarga) error {
	return r.db.Save(ashtakvarga).Error
}

func (r *AshtakvargaRepository) DeleteAshtakvarga(id uint) error {
	return r.db.Delete(&models.Ashtakvarga{}, id).Error
}
