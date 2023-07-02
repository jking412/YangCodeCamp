package model

type Type int

const (
	HTML Type = iota
	CSS
	JavaScript
)

type Question struct {
	ID          int    `gorm:"primaryKey"`
	Type        Type   `gorm:"column:type"`
	Answer      string `gorm:"column:answer"`
	Content     string `gorm:"column:content"`
	Description string `gorm:"column:description"`
}
