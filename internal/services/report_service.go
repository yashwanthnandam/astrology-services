package services

import (
	"astrology-services/internal/domain/models"
	"time"
)

type ReportService interface {
	GeneratePlanetaryReport(name, dob, tob, placeOfBirth string) (*models.PlanetaryReport, error)
	GenerateVimshottariReport(name, dob, tob, placeOfBirth string) (*models.VimshottariReport, error)
	GenerateYogaReport(name, dob, tob, placeOfBirth string) (*models.YogaReport, error)
}

type AstrologyReportService struct{}

func NewAstrologyReportService() ReportService {
	return &AstrologyReportService{}
}
func (s *AstrologyReportService) GeneratePlanetaryReport(name, dob, tob, placeOfBirth string) (*models.PlanetaryReport, error) {
	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}
	tobTime, err := time.Parse("15:04", tob)
	if err != nil {
		return nil, err
	}

	report := &PlanetaryReport{
		name: name,
		dob:  dobTime,
		tob:  tobTime,
		pob:  placeOfBirth,
	}

	return report.generateReport(), nil
}

func (s *AstrologyReportService) GenerateVimshottariReport(name, dob, tob, placeOfBirth string) (*models.VimshottariReport, error) {
	// Example implementation for Vimshottari report
	return &models.VimshottariReport{
		MercuryMahadasha: "Mercury Mahadasha: Placeholder details...",
		MoonMahadasha:    "Moon Mahadasha: Placeholder details...",
		VenusMahadasha:   "Venus Mahadasha: Placeholder details...",
		MarsMahadasha:    "Mars Mahadasha: Placeholder details...",
		JupiterMahadasha: "Jupiter Mahadasha: Placeholder details...",
		SaturnMahadasha:  "Saturn Mahadasha: Placeholder details...",
		RahuMahadasha:    "Rahu Mahadasha: Placeholder details...",
		KetuMahadasha:    "Ketu Mahadasha: Placeholder details...",
	}, nil
}

func (s *AstrologyReportService) GenerateYogaReport(name, dob, tob, placeOfBirth string) (*models.YogaReport, error) {
	// Example implementation for Yoga report
	return &models.YogaReport{
		GajakesariYoga:  "Gajakesari Yoga: Placeholder details...",
		SunaphaYoga:     "Sunapha Yoga: Placeholder details...",
		AdhiYoga:        "Adhi Yoga: Placeholder details...",
		BudhaAdityaYoga: "Budha-Aditya Yoga: Placeholder details...",
	}, nil
}
