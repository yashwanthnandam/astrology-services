package usecase

import "astrology-services/internal/domain/models"

type KundaliRepository interface {
	GetKundaliByID(id uint) (*models.Kundali, error)
}

type KundaliUsecase struct {
	repository KundaliRepository
}

func NewKundaliUsecase(r KundaliRepository) *KundaliUsecase {
	return &KundaliUsecase{r}
}

func (u *KundaliUsecase) GetKundali(id uint) (*models.Kundali, error) {
	return u.repository.GetKundaliByID(id)
}
