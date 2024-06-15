package http_test

// import (
// 	"astrology-services/internal/delivery/http"
// 	"astrology-services/internal/domain/models"
// 	"encoding/json"
// 	httpnet "net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockDoshaUsecase struct {
// 	mock.Mock
// }

// func (m *MockDoshaUsecase) GetDoshaAnalysis(id uint) (*models.DoshaAnalysis, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.DoshaAnalysis), args.Error(1)
// }

// func TestGetDoshaAnalysis(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockDoshaUsecase)
// 	analysis := &models.DoshaAnalysis{
// 		Mangalik: true,
// 		Kalsarpa: false,
// 		Sadesata: true,
// 	}

// 	mockUsecase.On("GetDoshaAnalysis", uint(1)).Return(analysis, nil)

// 	router := gin.Default()
// 	http.NewDoshaHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := httpnet.NewRequest("GET", "/dosha/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, httpnet.StatusOK, w.Code)

// 	var response models.DoshaAnalysis
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, analysis, &response)

// 	mockUsecase.AssertExpectations(t)
// }
