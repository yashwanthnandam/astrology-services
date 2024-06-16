package usecase

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/services"
)

type ReportRepository interface {
	GetGeneralReportByID(id uint) (*models.GeneralReport, error)
	GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error)
	GetVimshottariReportByID(id uint) (*models.VimshottariReport, error)
	GetYogaReportByID(id uint) (*models.YogaReport, error)
}

type ReportUsecase struct {
	repository    ReportRepository
	reportService services.ReportService
}

func NewReportUsecase(r ReportRepository, s services.ReportService) *ReportUsecase {
	return &ReportUsecase{repository: r, reportService: s}
}

func (u *ReportUsecase) GetGeneralReport(id uint) (*models.GeneralReport, error) {
	return u.repository.GetGeneralReportByID(id)
}

func (u *ReportUsecase) GeneratePlanetaryReport(name, dob, tob, placeOfBirth string) (*models.PlanetaryReport, error) {
	return u.reportService.GeneratePlanetaryReport(name, dob, tob, placeOfBirth)
}

func (u *ReportUsecase) GenerateVimshottariReport(name, dob, tob, placeOfBirth string) (*models.VimshottariReport, error) {
	return u.reportService.GenerateVimshottariReport(name, dob, tob, placeOfBirth)
}

func (u *ReportUsecase) GenerateYogaReport(name, dob, tob, placeOfBirth string) (*models.YogaReport, error) {
	return u.reportService.GenerateYogaReport(name, dob, tob, placeOfBirth)
}
