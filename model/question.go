package model

import (
	"gorm.io/gorm"
)

type Type int

const (
	HTML Type = iota
	CSS
	JavaScript
	PythonChoice
	Other
)

const (
	UnfinishedQuestion Status = iota
	PassQuestion
	FailQuestion
)

type Question struct {
	ID             int               `gorm:"primaryKey" json:"id"`
	Detail         []*QuestionDetail `json:"details"`
	Name           string            `json:"name"`
	Type           Type              `json:"type"`
	Status         Status            `json:"status"`
	NextQuestionID int               `json:"next_question_id"`
	PreQuestionID  int               `json:"pre_question_id"`
	ChapterID      int               `json:"chapter_id"`
}

type QuestionDetail struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Type         Type   `json:"type"`
	QuestionID   int    `json:"question_id"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	Answer       string `json:"answer"`
	CheckMessage string `json:"-"`
}

// AfterCreate
// 1. 在创建问题后更新章节的第一个问题和最后一个问题
// 2. 在创建问题后更新章节的问题总数
func (q *Question) AfterCreate(tx *gorm.DB) (err error) {
	if q.PreQuestionID == 0 {
		err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Update("first_question_id", q.ID).Error
		if err != nil {
			return err
		}
	}
	if q.NextQuestionID == 0 {
		err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Update("last_question_id", q.ID).Error
		if err != nil {
			return err
		}
	}

	err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Update("total_question", gorm.Expr("total_question + ?", 1)).Error
	if err != nil {
		return err
	}
	return
}

// AfterUpdate
// 1. 在更新问题后更新章节的问题总数
// 2. 在更新问题后更新章节的完成问题数
// 3. 在更新问题后更新章节的状态
func (q *Question) AfterUpdate(tx *gorm.DB) (err error) {
	questions := make([]*Question, 0)
	err = tx.Model(&Question{}).Where("chapter_id = ?", q.ChapterID).Find(&questions).Error
	if err != nil {
		return err
	}

	totalQuestion := len(questions)
	finishedQuestion := 0
	for _, question := range questions {
		if question.Status == PassQuestion {
			finishedQuestion++
		}
	}

	var status Status
	if finishedQuestion == totalQuestion {
		status = FinishedChapter
	} else {
		status = UnfinishedChapter
	}

	err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Updates(map[string]interface{}{
		"total_question":    totalQuestion,
		"finished_question": finishedQuestion,
		"status":            status,
	}).Error
	if err != nil {
		return err
	}

	return
}

func (q *Question) AfterDelete(tx *gorm.DB) (err error) {

	if q.PreQuestionID == 0 {
		err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Update("first_question_id", q.NextQuestionID).Error
		if err != nil {
			return err
		}

		if q.NextQuestionID != 0 {
			err = tx.Model(&Question{}).Where("id = ?", q.NextQuestionID).Update("pre_question_id", 0).Error
			if err != nil {
				return err
			}
		}

		return

	}

	if q.NextQuestionID == 0 {
		err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Update("last_question_id", q.PreQuestionID).Error
		if err != nil {
			return err
		}

		if q.PreQuestionID != 0 {
			err = tx.Model(&Question{}).Where("id = ?", q.PreQuestionID).Update("next_question_id", 0).Error
			if err != nil {
				return err
			}
		}

		return
	}

	questions := make([]*Question, 0)
	err = tx.Model(&Question{}).Where("chapter_id = ?", q.ChapterID).Find(&questions).Error
	if err != nil {
		return err
	}

	totalQuestion := len(questions)
	finishedQuestion := 0
	for _, question := range questions {
		if question.Status == PassQuestion {
			finishedQuestion++
		}
	}

	var status Status
	if finishedQuestion == totalQuestion {
		status = FinishedChapter
	} else {
		status = UnfinishedChapter
	}

	err = tx.Model(&Chapter{}).Where("id = ?", q.ChapterID).Updates(map[string]interface{}{
		"total_question":    totalQuestion,
		"finished_question": finishedQuestion,
		"status":            status,
	}).Error
	if err != nil {
		return err
	}

	return
}
