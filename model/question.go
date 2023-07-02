package model

type Question struct {
	ID          int    `gorm:"primaryKey"`
	Answer      string `gorm:"column:answer"`
	Content     string `gorm:"column:content"`
	Description string `gorm:"column:description"`
}
