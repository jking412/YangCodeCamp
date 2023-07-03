package model

type Type int

const (
	HTML Type = iota
	CSS
	JavaScript
)

type Question struct {
	ID             int `gorm:"primaryKey"`
	Type           Type
	Answer         string
	Content        string
	Description    string
	Name           string
	Number         int
	Status         Status
	NextQuestionID int
	PreQuestionID  int
	ChapterID      int
}
