package mysql

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/domain/repositories"

	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) repositories.ReportRepository {
	return &ReportRepository{db}
}

func (r *ReportRepository) GetGeneralReportByID(id uint) (*models.GeneralReport, error) {
	var report models.GeneralReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) CreateGeneralReport(report *models.GeneralReport) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) UpdateGeneralReport(report *models.GeneralReport) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) DeleteGeneralReport(id uint) error {
	return r.db.Delete(&models.GeneralReport{}, id).Error
}

func (r *ReportRepository) GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error) {
	var report models.PlanetaryReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) CreatePlanetaryReport(report *models.PlanetaryReport) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) UpdatePlanetaryReport(report *models.PlanetaryReport) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) DeletePlanetaryReport(id uint) error {
	return r.db.Delete(&models.PlanetaryReport{}, id).Error
}

func (r *ReportRepository) GetVimshottariReportByID(id uint) (*models.VimshottariReport, error) {
	var report models.VimshottariReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) CreateVimshottariReport(report *models.VimshottariReport) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) UpdateVimshottariReport(report *models.VimshottariReport) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) DeleteVimshottariReport(id uint) error {
	return r.db.Delete(&models.VimshottariReport{}, id).Error
}

func (r *ReportRepository) GetYogaReportByID(id uint) (*models.YogaReport, error) {
	var report models.YogaReport
	result := r.db.First(&report, id)
	return &report, result.Error
}

func (r *ReportRepository) CreateYogaReport(report *models.YogaReport) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) UpdateYogaReport(report *models.YogaReport) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) DeleteYogaReport(id uint) error {
	return r.db.Delete(&models.YogaReport{}, id).Error
}
