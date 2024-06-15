package http_test

// import (
// 	"astrology-services/internal/delivery/http"
// 	"astrology-services/internal/domain/models"
// 	"encoding/json"
// 	"fmt"
// 	httpnet "net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockDashaUsecase struct {
// 	mock.Mock
// }

// func (m *MockDashaUsecase) GetDasha(id uint) (*models.Dasha, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*models.Dasha), args.Error(1)
// }

// func TestGetDasha(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUsecase := new(MockDashaUsecase)
// 	fromDate, err := time.Parse("2006-01-02", "2020-01-01")
// 	if err != nil {
// 		fmt.Println("Error parsing FromDate:", err)
// 		return
// 	}

// 	toDate, err := time.Parse("2006-01-02", "2027-01-01")
// 	if err != nil {
// 		fmt.Println("Error parsing ToDate:", err)
// 		return
// 	}

// 	dasha := models.Dasha{
// 		ID:       1,
// 		FromDate: fromDate,
// 		ToDate:   toDate,
// 		Summary:  "This is a summary",
// 		Planet:   "Mars",
// 	}

// 	mockUsecase.On("GetDasha", uint(1)).Return(dasha, nil)

// 	router := gin.Default()
// 	http.NewDashaHandler(router, mockUsecase)

// 	w := httptest.NewRecorder()
// 	req, _ := httpnet.NewRequest("GET", "/dasha/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, httpnet.StatusOK, w.Code)

// 	var response models.Dasha
// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, dasha, &response)

// 	mockUsecase.AssertExpectations(t)
// }
