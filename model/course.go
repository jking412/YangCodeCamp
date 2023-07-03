package model

type Course struct {
	ID             int        `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"column:name;type:varchar(255)" json:"name"`
	Icon           string     `json:"icon"`
	Description    string     `gorm:"column:description;type:varchar(255)" json:"description"`
	FirstChapterID int        `json:"firstChapterID"`
	LastChapterID  int        `json:"lastChapterID"`
	Chapters       []*Chapter `gorm:"foreignKey:CourseID;references:ID" json:"-"`
}
