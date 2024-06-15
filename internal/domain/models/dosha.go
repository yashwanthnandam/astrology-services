package models

type DoshaAnalysis struct {
	Mangalik bool `gorm:"not null"`
	Kalsarpa bool `gorm:"not null"`
	Sadesata bool `gorm:"not null"`
}
