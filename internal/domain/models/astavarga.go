package models

type Ashtakvarga struct {
	ID      uint `gorm:"primaryKey"`
	Saturn  int  `gorm:"not null"`
	Sun     int  `gorm:"not null"`
	Venus   int  `gorm:"not null"`
	Mars    int  `gorm:"not null"`
	Mercury int  `gorm:"not null"`
	Moon    int  `gorm:"not null"`
	Sav     int  `gorm:"not null"`
	Asc     int  `gorm:"not null"`
	Jupiter int  `gorm:"not null"`
}
