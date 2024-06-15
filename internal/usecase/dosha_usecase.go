package usecase

import "astrology-services/internal/domain/models"

type DoshaRepository interface {
	GetDoshaAnalysisByID(id uint) (*models.DoshaAnalysis, error)
}

type DoshaUsecase struct {
	repository DoshaRepository
}

func NewDoshaUsecase(r DoshaRepository) *DoshaUsecase {
	return &DoshaUsecase{r}
}

func (u *DoshaUsecase) GetDoshaAnalysis(id uint) (*models.DoshaAnalysis, error) {
	return u.repository.GetDoshaAnalysisByID(id)
}
