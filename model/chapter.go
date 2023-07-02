package model

type Chapter struct {
	ID         int       `gorm:"primaryKey"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	Question   *Question `gorm:"foreignKey:QuestionID;references:ID"`
	ClassID    int
	QuestionID int
}
