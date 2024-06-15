package models

type PersonInfo struct {
	Name       string `gorm:"not null"`
	DOB        string `gorm:"not null"`
	BirthPlace string `gorm:"not null"`
}

type KundaliMatching struct {
	ID       uint       `gorm:"primaryKey"`
	GirlInfo PersonInfo `gorm:"embedded"`
	BoyInfo  PersonInfo `gorm:"embedded"`
}
