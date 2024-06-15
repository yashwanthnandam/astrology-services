package usecase_test

import (
	"astrology-services/internal/domain/models"
	"astrology-services/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRemedyRepository struct {
	mock.Mock
}

func (m *MockRemedyRepository) GetRudrakshaSuggestionByID(id uint) (*models.RudrakshaSuggestion, error) {
	args := m.Called(id)
	return args.Get(0).(*models.RudrakshaSuggestion), args.Error(1)
}

func (m *MockRemedyRepository) GetGemstoneSuggestionByID(id uint) (*models.GemstoneSuggestion, error) {
	args := m.Called(id)
	return args.Get(0).(*models.GemstoneSuggestion), args.Error(1)
}

func TestGetRudrakshaSuggestion(t *testing.T) {
	mockRepo := new(MockRemedyRepository)
	suggestion := &models.RudrakshaSuggestion{
		Recommendation: "Wear 5 Mukhi Rudraksha",
		Details:        "Details about 5 Mukhi Rudraksha",
		Benefit:        "Benefits of wearing 5 Mukhi Rudraksha",
		HowToWear:      "How to wear 5 Mukhi Rudraksha",
		Precautions:    "Precautions while wearing 5 Mukhi Rudraksha",
	}

	mockRepo.On("GetRudrakshaSuggestionByID", uint(1)).Return(suggestion, nil)

	remedyUsecase := usecase.NewRemedyUsecase(mockRepo)
	result, err := remedyUsecase.GetRudrakshaSuggestion(1)

	assert.NoError(t, err)
	assert.Equal(t, suggestion, result)

	mockRepo.AssertExpectations(t)
}

func TestGetGemstoneSuggestion(t *testing.T) {
	mockRepo := new(MockRemedyRepository)
	suggestion := &models.GemstoneSuggestion{
		LifeStone:    "Ruby",
		LuckyStone:   "Emerald",
		FortuneStone: "Diamond",
	}

	mockRepo.On("GetGemstoneSuggestionByID", uint(1)).Return(suggestion, nil)

	remedyUsecase := usecase.NewRemedyUsecase(mockRepo)
	result, err := remedyUsecase.GetGemstoneSuggestion(1)

	assert.NoError(t, err)
	assert.Equal(t, suggestion, result)

	mockRepo.AssertExpectations(t)
}
