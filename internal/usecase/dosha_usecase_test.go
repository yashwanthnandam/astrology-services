package usecase_test

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDoshaRepository struct {
	mock.Mock
}

func (m *MockDoshaRepository) GetDoshaAnalysisByID(id uint) (*models.DoshaAnalysis, error) {
	args := m.Called(id)
	return args.Get(0).(*models.DoshaAnalysis), args.Error(1)
}

func TestGetDoshaAnalysis(t *testing.T) {
	mockRepo := new(MockDoshaRepository)
	analysis := &models.DoshaAnalysis{
		Mangalik: true,
		Kalsarpa: false,
		Sadesata: true,
	}

	mockRepo.On("GetDoshaAnalysisByID", uint(1)).Return(analysis, nil)

	doshaUsecase := usecase.NewDoshaUsecase(mockRepo)
	result, err := doshaUsecase.GetDoshaAnalysis(1)

	assert.NoError(t, err)
	assert.Equal(t, analysis, result)

	mockRepo.AssertExpectations(t)
}
