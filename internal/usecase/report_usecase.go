package usecase

import "astrology-services/internal/domain/models"

type ReportRepository interface {
	GetGeneralReportByID(id uint) (*models.GeneralReport, error)
	GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error)
	GetVimshottariReportByID(id uint) (*models.VimshottariReport, error)
	GetYogaReportByID(id uint) (*models.YogaReport, error)
}

type ReportUsecase struct {
	repository ReportRepository
}

func NewReportUsecase(r ReportRepository) *ReportUsecase {
	return &ReportUsecase{r}
}

func (u *ReportUsecase) GetGeneralReport(id uint) (*models.GeneralReport, error) {
	return u.repository.GetGeneralReportByID(id)
}

func (u *ReportUsecase) GetPlanetaryReport(id uint) (*models.PlanetaryReport, error) {
	return u.repository.GetPlanetaryReportByID(id)
}

func (u *ReportUsecase) GetVimshottariReport(id uint) (*models.VimshottariReport, error) {
	return u.repository.GetVimshottariReportByID(id)
}

func (u *ReportUsecase) GetYogaReport(id uint) (*models.YogaReport, error) {
	return u.repository.GetYogaReportByID(id)
}
