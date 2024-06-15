package usecase_test

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReportRepository struct {
	mock.Mock
}

func (m *MockReportRepository) GetGeneralReportByID(id uint) (*models.GeneralReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.GeneralReport), args.Error(1)
}

func (m *MockReportRepository) GetPlanetaryReportByID(id uint) (*models.PlanetaryReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.PlanetaryReport), args.Error(1)
}

func (m *MockReportRepository) GetVimshottariReportByID(id uint) (*models.VimshottariReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.VimshottariReport), args.Error(1)
}

func (m *MockReportRepository) GetYogaReportByID(id uint) (*models.YogaReport, error) {
	args := m.Called(id)
	return args.Get(0).(*models.YogaReport), args.Error(1)
}

func TestGetGeneralReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	report := &models.GeneralReport{
		Description:  "This is a description",
		Personality:  "This is a personality",
		Physical:     "This is physical",
		Health:       "This is health",
		Career:       "This is career",
		Relationship: "This is relationship",
	}

	mockRepo.On("GetGeneralReportByID", uint(1)).Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo)
	result, err := reportUsecase.GetGeneralReport(1)

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockRepo.AssertExpectations(t)
}

func TestGetPlanetaryReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	report := &models.PlanetaryReport{
		SunConsideration:     "Sun is strong",
		MoonConsideration:    "Moon is strong",
		MercuryConsideration: "Mercury is strong",
		VenusConsideration:   "Venus is strong",
		MarsConsideration:    "Mars is strong",
		JupiterConsideration: "Jupiter is strong",
		SaturnConsideration:  "Saturn is strong",
		RahuConsideration:    "Rahu is strong",
		KetuConsideration:    "Ketu is strong",
	}

	mockRepo.On("GetPlanetaryReportByID", uint(1)).Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo)
	result, err := reportUsecase.GetPlanetaryReport(1)

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockRepo.AssertExpectations(t)
}

func TestGetVimshottariReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	report := &models.VimshottariReport{
		MercuryMahadasha: "Mercury Mahadasha is good",
		MoonMahadasha:    "Moon Mahadasha is good",
		VenusMahadasha:   "Venus Mahadasha is good",
		MarsMahadasha:    "Mars Mahadasha is good",
		JupiterMahadasha: "Jupiter Mahadasha is good",
		SaturnMahadasha:  "Saturn Mahadasha is good",
		RahuMahadasha:    "Rahu Mahadasha is good",
		KetuMahadasha:    "Ketu Mahadasha is good",
	}

	mockRepo.On("GetVimshottariReportByID", uint(1)).Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo)
	result, err := reportUsecase.GetVimshottariReport(1)

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockRepo.AssertExpectations(t)
}

func TestGetYogaReport(t *testing.T) {
	mockRepo := new(MockReportRepository)
	report := &models.YogaReport{
		GajakesariYoga:  "Gajakesari Yoga is present",
		SunaphaYoga:     "Sunapha Yoga is present",
		AdhiYoga:        "Adhi Yoga is present",
		BudhaAdityaYoga: "Budha-Aditya Yoga is present",
	}

	mockRepo.On("GetYogaReportByID", uint(1)).Return(report, nil)

	reportUsecase := usecase.NewReportUsecase(mockRepo)
	result, err := reportUsecase.GetYogaReport(1)

	assert.NoError(t, err)
	assert.Equal(t, report, result)

	mockRepo.AssertExpectations(t)
}
