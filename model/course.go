package model

const (
	UnfinishedCourse Status = iota
	FinishedCourse
)

type Course struct {
	ID          int        `gorm:"primaryKey"`
	Name        string     `gorm:"column:name;type:varchar(255)"`
	Description string     `gorm:"column:description;type:varchar(255)"`
	Chapters    []*Chapter `gorm:"foreignKey:CourseID;references:ID"`
	Status      Status
}
