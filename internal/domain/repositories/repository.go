package repositories

import "astrology-services/internal/domain/models"

// KundaliRepository defines methods for accessing Kundali data
type KundaliRepository interface {
	GetKundaliByID(id uint) (*models.Kundali, error)
	CreateKundali(kundali *models.Kundali) error
	UpdateKundali(kundali *models.Kundali) error
	DeleteKundali(id uint) error
}

// AshtakvargaRepository defines methods for accessing Ashtakvarga data
type AshtakvargaRepository interface {
	GetAshtakvargaByID(id uint) (*models.Ashtakvarga, error)
	CreateAshtakvarga(ashtakvarga *models.Ashtakvarga) error
	UpdateAshtakvarga(ashtakvarga *models.Ashtakvarga) error
	DeleteAshtakvarga(id uint) error
}

// DashaRepository defines methods for accessing Dasha data
type DashaRepository interface {
	GetDashaByID(id uint) (*models.Dasha, error)
	CreateDasha(dasha *models.Dasha) error
	UpdateDasha(dasha *models.Dasha) error
	DeleteDasha(id uint) error
}

// ReportRepository defines methods for accessing Report data
type ReportRepository interface {
	GetGeneralReportByID(id uint) (*models.GeneralReport, error)
	CreateGeneralReport(report *models.GeneralReport) error
	UpdateGeneralReport(report *models.GeneralReport) error
	DeleteGeneralReport(id uint) error

	GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error)
	CreatePlanetaryReport(report *models.PlanetaryReport) error
	UpdatePlanetaryReport(report *models.PlanetaryReport) error
	DeletePlanetaryReport(id uint) error

	GetVimshottariReportByID(id uint) (*models.VimshottariReport, error)
	CreateVimshottariReport(report *models.VimshottariReport) error
	UpdateVimshottariReport(report *models.VimshottariReport) error
	DeleteVimshottariReport(id uint) error

	GetYogaReportByID(id uint) (*models.YogaReport, error)
	CreateYogaReport(report *models.YogaReport) error
	UpdateYogaReport(report *models.YogaReport) error
	DeleteYogaReport(id uint) error
}

// RemedyRepository defines methods for accessing Remedy data
type RemedyRepository interface {
	GetRudrakshaSuggestionByID(id uint) (*models.RudrakshaSuggestion, error)
	CreateRudrakshaSuggestion(suggestion *models.RudrakshaSuggestion) error
	UpdateRudrakshaSuggestion(suggestion *models.RudrakshaSuggestion) error
	DeleteRudrakshaSuggestion(id uint) error

	GetGemstoneSuggestionByID(id uint) (*models.GemstoneSuggestion, error)
	CreateGemstoneSuggestion(suggestion *models.GemstoneSuggestion) error
	UpdateGemstoneSuggestion(suggestion *models.GemstoneSuggestion) error
	DeleteGemstoneSuggestion(id uint) error
}

// DoshaRepository defines methods for accessing Dosha data
type DoshaRepository interface {
	GetDoshaAnalysisByID(id uint) (*models.DoshaAnalysis, error)
	CreateDoshaAnalysis(analysis *models.DoshaAnalysis) error
	UpdateDoshaAnalysis(analysis *models.DoshaAnalysis) error
	DeleteDoshaAnalysis(id uint) error
}

// MatchRepository defines methods for accessing Kundali Matching data
type MatchRepository interface {
	GetKundaliMatchingByID(id uint) (*models.KundaliMatching, error)
	CreateKundaliMatching(matching *models.KundaliMatching) error
	UpdateKundaliMatching(matching *models.KundaliMatching) error
	DeleteKundaliMatching(id uint) error
}
