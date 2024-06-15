package usecase_test

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMatchRepository struct {
	mock.Mock
}

func (m *MockMatchRepository) GetKundaliMatchingByID(id uint) (*models.KundaliMatching, error) {
	args := m.Called(id)
	return args.Get(0).(*models.KundaliMatching), args.Error(1)
}

func TestGetKundaliMatching(t *testing.T) {
	mockRepo := new(MockMatchRepository)
	matching := &models.KundaliMatching{
		ID: 1,
		GirlInfo: models.PersonInfo{
			Name:       "Alice",
			DOB:        "1990-01-01",
			BirthPlace: "New York",
		},
		BoyInfo: models.PersonInfo{
			Name:       "Bob",
			DOB:        "1990-01-01",
			BirthPlace: "Los Angeles",
		},
	}

	mockRepo.On("GetKundaliMatchingByID", uint(1)).Return(matching, nil)

	matchUsecase := usecase.NewMatchUsecase(mockRepo)
	result, err := matchUsecase.GetKundaliMatching(1)

	assert.NoError(t, err)
	assert.Equal(t, matching, result)

	mockRepo.AssertExpectations(t)
}
