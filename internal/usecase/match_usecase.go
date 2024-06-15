package usecase

import "astrology-services/internal/domain/models"

type MatchRepository interface {
	GetKundaliMatchingByID(id uint) (*models.KundaliMatching, error)
}

type MatchUsecase struct {
	repository MatchRepository
}

func NewMatchUsecase(r MatchRepository) *MatchUsecase {
	return &MatchUsecase{r}
}

func (u *MatchUsecase) GetKundaliMatching(id uint) (*models.KundaliMatching, error) {
	return u.repository.GetKundaliMatchingByID(id)
}
