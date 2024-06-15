package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type RemedyRepository struct {
	db *gorm.DB
}

func NewRemedyRepository(db *gorm.DB) repositories.RemedyRepository {
	return &RemedyRepository{db}
}

func (r *RemedyRepository) GetRudrakshaSuggestionByID(id uint) (*models.RudrakshaSuggestion, error) {
	var suggestion models.RudrakshaSuggestion
	result := r.db.First(&suggestion, id)
	return &suggestion, result.Error
}

func (r *RemedyRepository) CreateRudrakshaSuggestion(suggestion *models.RudrakshaSuggestion) error {
	return r.db.Create(suggestion).Error
}

func (r *RemedyRepository) UpdateRudrakshaSuggestion(suggestion *models.RudrakshaSuggestion) error {
	return r.db.Save(suggestion).Error
}

func (r *RemedyRepository) DeleteRudrakshaSuggestion(id uint) error {
	return r.db.Delete(&models.RudrakshaSuggestion{}, id).Error
}

func (r *RemedyRepository) GetGemstoneSuggestionByID(id uint) (*models.GemstoneSuggestion, error) {
	var suggestion models.GemstoneSuggestion
	result := r.db.First(&suggestion, id)
	return &suggestion, result.Error
}

func (r *RemedyRepository) CreateGemstoneSuggestion(suggestion *models.GemstoneSuggestion) error {
	return r.db.Save(suggestion).Error
}

func (r *RemedyRepository) UpdateGemstoneSuggestion(suggestion *models.GemstoneSuggestion) error {
	return r.db.Save(suggestion).Error
}
func (r *RemedyRepository) DeleteGemstoneSuggestion(id uint) error {
	return r.db.Delete(&models.RudrakshaSuggestion{}, id).Error
}
