package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) repositories.MatchRepository {
	return &MatchRepository{db}
}

func (r *MatchRepository) GetKundaliMatchingByID(id uint) (*models.KundaliMatching, error) {
	var matching models.KundaliMatching
	result := r.db.First(&matching, id)
	return &matching, result.Error
}

func (r *MatchRepository) CreateKundaliMatching(matching *models.KundaliMatching) error {
	return r.db.Create(matching).Error
}

func (r *MatchRepository) UpdateKundaliMatching(matching *models.KundaliMatching) error {
	return r.db.Save(matching).Error
}

func (r *MatchRepository) DeleteKundaliMatching(id uint) error {
	return r.db.Delete(&models.KundaliMatching{}, id).Error
}
