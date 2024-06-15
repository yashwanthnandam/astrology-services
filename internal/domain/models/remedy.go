package models

type RudrakshaSuggestion struct {
	Recommendation string `gorm:"not null"`
	Details        string `gorm:"not null"`
	Benefit        string `gorm:"not null"`
	HowToWear      string `gorm:"not null"`
	Precautions    string `gorm:"not null"`
}

type GemstoneSuggestion struct {
	LifeStone    string `gorm:"not null"`
	LuckyStone   string `gorm:"not null"`
	FortuneStone string `gorm:"not null"`
}
