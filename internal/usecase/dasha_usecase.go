package usecase

import "astrology-services/internal/domain/models"

type DashaRepository interface {
	GetDashaByID(id uint) (*models.Dasha, error)
}

type DashaUsecase struct {
	repository DashaRepository
}

func NewDashaUsecase(r DashaRepository) *DashaUsecase {
	return &DashaUsecase{r}
}

func (u *DashaUsecase) GetDasha(id uint) (*models.Dasha, error) {
	return u.repository.GetDashaByID(id)
}
