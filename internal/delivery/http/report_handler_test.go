package http_test

// import (
// 	"astrology-services/internal/domain/models"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockReportUsecase struct {
// 	mock.Mock
// }

// func (m *MockReportUsecase) GetGeneralReport(id uint) (*models.GeneralReport, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.GeneralReport), args.Error(1)
// }

// func (m *MockReportUsecase) GetPlanetaryReport(id uint) (*models.PlanetaryReport, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.PlanetaryReport), args.Error(1)
// }

// func (m *MockReportUsecase) GetVimshottariReport(id uint) (*models.VimshottariReport, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.VimshottariReport), args.Error(1)
// }

// func (m *MockReportUsecase) GetYogaReport(id uint) (*models.YogaReport, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.YogaReport), args.Error(1)
// }

// func TestGetGeneralReport(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockReportUsecase)
// 	report := &models.GeneralReport{
// 		Description:  "This is a description",
// 		Personality:  "This is a personality",
// 		Physical:     "This is physical",
// 		Health:       "This is health",
// 		Career:       "This is career",
// 		Relationship: "This is relationship",
// 	}

// 	mockUsecase.On("GetGeneralReport", uint(1)).Return(report, nil)

// 	router := gin.Default()
// 	http.NewReportHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/report/general/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.GeneralReport
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, report, &response)

// 	mockUsecase.AssertExpectations(t)
// }

// func TestGetPlanetaryReport(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockReportUsecase)
// 	report := &models.PlanetaryReport{
// 		SunConsideration:     "Sun is strong",
// 		MoonConsideration:    "Moon is strong",
// 		MercuryConsideration: "Mercury is strong",
// 		VenusConsideration:   "Venus is strong",
// 		MarsConsideration:    "Mars is strong",
// 		JupiterConsideration: "Jupiter is strong",
// 		SaturnConsideration:  "Saturn is strong",
// 		RahuConsideration:    "Rahu is strong",
// 		KetuConsideration:    "Ketu is strong",
// 	}

// 	mockUsecase.On("GetPlanetaryReport", uint(1)).Return(report, nil)

// 	router := gin.Default()
// 	http.NewReportHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/report/planetary/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.PlanetaryReport
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, report, &response)

// 	mockUsecase.AssertExpectations(t)
// }

// func TestGetVimshottariReport(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockReportUsecase)
// 	report := &models.VimshottariReport{
// 		MercuryMahadasha: "Mercury Mahadasha is good",
// 		MoonMahadasha:    "Moon Mahadasha is good",
// 		VenusMahadasha:   "Venus Mahadasha is good",
// 		MarsMahadasha:    "Mars Mahadasha is good",
// 		JupiterMahadasha: "Jupiter Mahadasha is good",
// 		SaturnMahadasha:  "Saturn Mahadasha is good",
// 		RahuMahadasha:    "Rahu Mahadasha is good",
// 		KetuMahadasha:    "Ketu Mahadasha is good",
// 	}

// 	mockUsecase.On("GetVimshottariReport", uint(1)).Return(report, nil)

// 	router := gin.Default()
// 	http.NewReportHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/report/vimshottari/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.VimshottariReport
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, report, &response)

// 	mockUsecase.AssertExpectations(t)
// }

// func TestGetYogaReport(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockReportUsecase)
// 	report := &models.YogaReport{
// 		GajakesariYoga:  "Gajakesari Yoga is present",
// 		SunaphaYoga:     "Sunapha Yoga is present",
// 		AdhiYoga:        "Adhi Yoga is present",
// 		BudhaAdityaYoga: "Budha-Aditya Yoga is present",
// 	}

// 	mockUsecase.On("GetYogaReport", uint(1)).Return(report, nil)

// 	router := gin.Default()
// 	http.NewReportHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/report/yoga/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.YogaReport
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, report, &response)

// 	mockUsecase.AssertExpectations(t)
// }
