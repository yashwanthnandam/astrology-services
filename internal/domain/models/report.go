package models

type GeneralReport struct {
	Description  string `gorm:"not null"`
	Personality  string `gorm:"not null"`
	Physical     string `gorm:"not null"`
	Health       string `gorm:"not null"`
	Career       string `gorm:"not null"`
	Relationship string `gorm:"not null"`
}

type PlanetaryReport struct {
	SunConsideration     string `gorm:"not null"`
	MoonConsideration    string `gorm:"not null"`
	MercuryConsideration string `gorm:"not null"`
	VenusConsideration   string `gorm:"not null"`
	MarsConsideration    string `gorm:"not null"`
	JupiterConsideration string `gorm:"not null"`
	SaturnConsideration  string `gorm:"not null"`
	RahuConsideration    string `gorm:"not null"`
	KetuConsideration    string `gorm:"not null"`
}

type VimshottariReport struct {
	MercuryMahadasha string `gorm:"not null"`
	MoonMahadasha    string `gorm:"not null"`
	VenusMahadasha   string `gorm:"not null"`
	MarsMahadasha    string `gorm:"not null"`
	JupiterMahadasha string `gorm:"not null"`
	SaturnMahadasha  string `gorm:"not null"`
	RahuMahadasha    string `gorm:"not null"`
	KetuMahadasha    string `gorm:"not null"`
}

type YogaReport struct {
	GajakesariYoga  string `gorm:"not null"`
	SunaphaYoga     string `gorm:"not null"`
	AdhiYoga        string `gorm:"not null"`
	BudhaAdityaYoga string `gorm:"not null"`
}
