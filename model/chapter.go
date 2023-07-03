package model

import "gorm.io/gorm"

type Status int

const (
	UnfinishedChapter Status = iota
	FinishedChapter
)

type Chapter struct {
	ID               int    `gorm:"primaryKey"`
	Name             string `gorm:"column:name;type:varchar(255)"`
	Description      string
	Status           Status
	PreChapterID     int
	NextChapterID    int
	FirstQuestionID  int
	LastQuestionID   int
	TotalQuestion    int
	FinishedQuestion int
	CourseID         int
	Question         []*Question `gorm:"foreignKey:ChapterID;references:ID" json:"-"`
}

// AfterCreate
// 1. 在创建章节后更新课程的第一个章节和最后一个章节
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

// AfterDelete
// 1. 在删除章节后更新课程的第一个章节和最后一个章节
// 2. 在删除章节后更新章节的前一个章节和后一个章节
func (c *Chapter) AfterDelete(tx *gorm.DB) (err error) {
	if c.PreChapterID == 0 {
		err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("first_chapter_id", c.NextChapterID).Error
		if err != nil {
			return err
		}

		if c.NextChapterID != 0 {
			err = tx.Model(&Chapter{}).Where("id = ?", c.NextChapterID).Update("pre_chapter_id", 0).Error
			if err != nil {
				return err
			}
		}

	}
	if c.NextChapterID == 0 {
		err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("last_chapter_id", c.PreChapterID).Error
		if err != nil {
			return err
		}

		if c.PreChapterID != 0 {
			err = tx.Model(&Chapter{}).Where("id = ?", c.PreChapterID).Update("next_chapter_id", 0).Error
			if err != nil {
				return err
			}
		}
	}
	return
}

//
//func (c *Chapter) AfterUpdate(tx *gorm.DB) (err error) {
//	if c.Status == SuccessChapter {
//		var count int64
//		err = tx.Model(&Chapter{}).Where("status in (?) and course_id = ?",
//			[]Status{
//				UnfinishedChapter,
//				FailChapter,
//			},
//			c.CourseID).Count(&count).Error
//		if err != nil {
//			return err
//		}
//		if count == 0 {
//			err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("status", FinishedCourse).Error
//			if err != nil {
//				return err
//			}
//		}
//		return
//	}
//
//	err = tx.Model(&Course{}).Where("id = ?", c.CourseID).Update("status", UnfinishedCourse).Error
//	if err != nil {
//		return err
//	}
//
//	return
//}
