package mysql

import (
	"astrology-services/internal/domain/models"

	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db}
}

func (r *ReportRepository) GetGeneralReportByID(id uint) (*models.GeneralReport, error) {
	var report models.GeneralReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error) {
	var report models.PlanetaryReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) GetVimshottariReportByID(id uint) (*models.VimshottariReport, error) {
	var report models.VimshottariReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) GetYogaReportByID(id uint) (*models.YogaReport, error) {
	var report models.YogaReport
	result := r.db.First(&report, id)
	return &report, result.Error
}
