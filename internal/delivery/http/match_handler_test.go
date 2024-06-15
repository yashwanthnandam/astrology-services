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

// type MockMatchUsecase struct {
// 	mock.Mock
// }

// func (m *MockMatchUsecase) GetKundaliMatching(id uint) (*models.KundaliMatching, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.KundaliMatching), args.Error(1)
// }

// func TestGetKundaliMatching(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockMatchUsecase)
// 	matching := &models.KundaliMatching{
// 		ID: 1,
// 		GirlInfo: models.PersonInfo{
// 			Name:       "Alice",
// 			DOB:        "1990-01-01",
// 			BirthPlace: "New York",
// 		},
// 		BoyInfo: models.PersonInfo{
// 			Name:       "Bob",
// 			DOB:        "1990-01-01",
// 			BirthPlace: "Los Angeles",
// 		},
// 	}

// 	mockUsecase.On("GetKundaliMatching", uint(1)).Return(matching, nil)

// 	router := gin.Default()
// 	http.NewMatchHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/match/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.KundaliMatching
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, matching, &response)

// 	mockUsecase.AssertExpectations(t)
// }
