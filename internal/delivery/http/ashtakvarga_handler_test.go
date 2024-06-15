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

// type MockAshtakvargaUsecase struct {
// 	mock.Mock
// }

// func (m *MockAshtakvargaUsecase) GetAshtakvarga(id uint) (*models.Ashtakvarga, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.Ashtakvarga), args.Error(1)
// }

// func TestGetAshtakvarga(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockAshtakvargaUsecase)
// 	ashtakvarga := &models.Ashtakvarga{
// 		ID:      1,
// 		Saturn:  2,
// 		Sun:     3,
// 		Venus:   4,
// 		Mars:    5,
// 		Mercury: 6,
// 		Moon:    7,
// 		Sav:     8,
// 		Asc:     9,
// 		Jupiter: 10,
// 	}

// 	mockUsecase.On("GetAshtakvarga", uint(1)).Return(ashtakvarga, nil)

// 	router := gin.Default()
// 	http.NewAshtakvargaHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := httpnet.NewRequest("GET", "/ashtakvarga/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, httpnet.StatusOK, w.Code)

// 	var response models.Ashtakvarga
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, ashtakvarga, &response)

// 	mockUsecase.AssertExpectations(t)
// }
