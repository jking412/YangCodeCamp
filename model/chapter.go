package model

type Chapter struct {
	ID        int         `gorm:"primaryKey"`
	Name      string      `gorm:"column:name;type:varchar(255)"`
	Questions []*Question `gorm:"foreignKey:ChapterID;references:ID"`
	ClassID   int
}
