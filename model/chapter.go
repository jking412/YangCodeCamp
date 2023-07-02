package model

import "gorm.io/gorm"

type Status int

const (
	UnfinishedChapter Status = iota
	SuccessChapter
	FailChapter
)

type Chapter struct {
	ID            int `gorm:"primaryKey"`
	Number        int
	Name          string    `gorm:"column:name;type:varchar(255)"`
	Question      *Question `gorm:"foreignKey:QuestionID;references:ID"`
	Status        Status
	PreChapterID  int
	NextChapterID int
	CourseID      int
	QuestionID    int
}

func (c *Chapter) AfterCreate(tx *gorm.DB) (err error) {
	if c.PreChapterID == 0 {
		err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("first_chapter_id", c.ID).Error
		if err != nil {
			return err
		}
		return
	}
	if c.NextChapterID == 0 {
		err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("last_chapter_id", c.ID).Error
		if err != nil {
			return err
		}
		return
	}
	return
}

func (c *Chapter) AfterUpdate(tx *gorm.DB) (err error) {
	if c.Status == SuccessChapter {
		var count int64
		err = tx.Model(&Chapter{}).Where("status in (?) and course_id = ?",
			[]Status{
				UnfinishedChapter,
				FailChapter,
			},
			c.CourseID).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("status", FinishedCourse).Error
			if err != nil {
				return err
			}
		}
		return
	}

	err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("status", UnfinishedCourse).Error
	if err != nil {
		return err
	}

	return
}
