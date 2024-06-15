package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type DoshaRepository struct {
	db *gorm.DB
}

func NewDoshaRepository(db *gorm.DB) repositories.DoshaRepository {
	return &DoshaRepository{db}
}

func (r *DoshaRepository) GetDoshaAnalysisByID(id uint) (*models.DoshaAnalysis, error) {
	var analysis models.DoshaAnalysis
	result := r.db.First(&analysis, id)
	return &analysis, result.Error
}

func (r *DoshaRepository) CreateDoshaAnalysis(analysis *models.DoshaAnalysis) error {
	return r.db.Create(analysis).Error
}

func (r *DoshaRepository) UpdateDoshaAnalysis(analysis *models.DoshaAnalysis) error {
	return r.db.Save(analysis).Error
}

func (r *DoshaRepository) DeleteDoshaAnalysis(id uint) error {
	return r.db.Delete(&models.DoshaAnalysis{}, id).Error
}
