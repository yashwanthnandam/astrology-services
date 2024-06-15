package models

type Kundali struct {
	ID        uint    `gorm:"primaryKey"`
	Planet    string  `gorm:"not null"`
	Sign      string  `gorm:"not null"`
	SignLord  string  `gorm:"not null"`
	Nakshatra string  `gorm:"not null"`
	NakshLord string  `gorm:"not null"`
	Degree    float64 `gorm:"not null"`
	Retro     bool    `gorm:"not null"`
	Combust   bool    `gorm:"not null"`
	Avastha   string  `gorm:"not null"`
	House     int     `gorm:"not null"`
	Status    string  `gorm:"not null"`
}
