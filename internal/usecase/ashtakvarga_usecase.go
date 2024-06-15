package usecase

import "astrology-services/internal/domain/models"

type AshtakvargaRepository interface {
	GetAshtakvargaByID(id uint) (*models.Ashtakvarga, error)
}

type AshtakvargaUsecase struct {
	repository AshtakvargaRepository
}

func NewAshtakvargaUsecase(r AshtakvargaRepository) *AshtakvargaUsecase {
	return &AshtakvargaUsecase{r}
}

func (u *AshtakvargaUsecase) GetAshtakvarga(id uint) (*models.Ashtakvarga, error) {
	return u.repository.GetAshtakvargaByID(id)
}
