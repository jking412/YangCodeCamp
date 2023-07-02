package model

type Status int

const (
	UnfinishedChapter Status = iota
	SuccessChapter
	FailChapter
)

type Chapter struct {
	ID         int       `gorm:"primaryKey"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	Question   *Question `gorm:"foreignKey:QuestionID;references:ID"`
	Status     Status
	CourseID   int
	QuestionID int
}
