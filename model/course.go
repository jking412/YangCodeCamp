package model

type Course struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"column:name;type:varchar(255)"`
	Icon           string
	Description    string `gorm:"column:description;type:varchar(255)"`
	FirstChapterID int
	LastChapterID  int
	Chapters       []*Chapter `gorm:"foreignKey:CourseID;references:ID" json:"-"`
}
