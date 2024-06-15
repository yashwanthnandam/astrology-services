package http

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

// type MockKundaliUsecase struct {
// 	mock.Mock
// }

// func (m *MockKundaliUsecase) GetKundali(id uint) (*models.Kundali, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.Kundali), args.Error(1)
// }

// func TestGetKundali(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockKundaliUsecase)
// 	kundali := &models.Kundali{
// 		ID:        1,
// 		Planet:    "Mars",
// 		Sign:      "Aries",
// 		SignLord:  "Mars",
// 		Nakshatra: "Ashwini",
// 		NakshLord: "Ketu",
// 		Degree:    1.0,
// 		Retro:     false,
// 		Combust:   false,
// 		Avastha:   "Awake",
// 		House:     1,
// 		Status:    "Good",
// 	}

// 	mockUsecase.On("GetKundali", uint(1)).Return(kundali, nil)

// 	router := gin.Default()
// 	NewKundaliHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/kundali/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.Kundali
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, kundali, &response)

// 	mockUsecase.AssertExpectations(t)
// }
