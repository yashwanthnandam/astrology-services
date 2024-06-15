package usecase

import "astrology-services/internal/domain/models"

type RemedyRepository interface {
	GetRudrakshaSuggestionByID(id uint) (*models.RudrakshaSuggestion, error)
	GetGemstoneSuggestionByID(id uint) (*models.GemstoneSuggestion, error)
}

type RemedyUsecase struct {
	repository RemedyRepository
}

func NewRemedyUsecase(r RemedyRepository) *RemedyUsecase {
	return &RemedyUsecase{r}
}

func (u *RemedyUsecase) GetRudrakshaSuggestion(id uint) (*models.RudrakshaSuggestion, error) {
	return u.repository.GetRudrakshaSuggestionByID(id)
}

func (u *RemedyUsecase) GetGemstoneSuggestion(id uint) (*models.GemstoneSuggestion, error) {
	return u.repository.GetGemstoneSuggestionByID(id)
}
