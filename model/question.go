package model

type Question struct {
	ID        int    `gorm:"primaryKey"`
	Answer    string `gorm:"column:answer"`
	ChapterID int
}
