package http_test

// import (
// 	"astrology-services/internal/domain/models"
// 	"astrology-services/internal/usecase"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRemedyUsecase struct {
// 	mock.Mock
// }

// func (m *MockRemedyUsecase) GetRudrakshaSuggestion(id uint) (*models.RudrakshaSuggestion, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.RudrakshaSuggestion), args.Error(1)
// }

// func (m *MockRemedyUsecase) GetGemstoneSuggestion(id uint) (*models.GemstoneSuggestion, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.GemstoneSuggestion), args.Error(1)
// }

// func TestGetRudrakshaSuggestion(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockRemedyUsecase)
// 	suggestion := &models.RudrakshaSuggestion{
// 		Recommendation: "Wear 5 Mukhi Rudraksha",
// 		Details:        "Details about 5 Mukhi Rudraksha",
// 		Benefit:        "Benefits of wearing 5 Mukhi Rudraksha",
// 		HowToWear:      "How to wear 5 Mukhi Rudraksha",
// 		Precautions:    "Precautions while wearing 5 Mukhi Rudraksha",
// 	}

// 	mockUsecase.On("GetRudrakshaSuggestion", uint(1)).Return(suggestion, nil)

// 	router := gin.Default()
// 	http.NewRemedyHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/remedy/rudraksha/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.RudrakshaSuggestion
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, suggestion, &response)

// 	mockUsecase.AssertExpectations(t)
// }

// func TestGetGemstoneSuggestion(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockRemedyUsecase)
// 	suggestion := &models.GemstoneSuggestion{
// 		LifeStone:    "Ruby",
// 		LuckyStone:   "Emerald",
// 		FortuneStone: "Diamond",
// 	}

// 	mockUsecase.On("GetGemstoneSuggestion", uint(1)).Return(suggestion, nil)

// 	router := gin.Default()
// 	http.NewRemedyHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.
