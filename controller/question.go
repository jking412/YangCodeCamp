package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/answers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetQuestionById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	question := &model.Question{}
	err = db.Mysql.Model(&model.Question{}).Where("id = ?", id).First(&question).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, question)
}

type SubmitQuestionReq struct {
	Content string `json:"content"`
}

type SubmitQuestionResp struct {
}

func SubmitQuestion(c *gin.Context) {

	req := &SubmitQuestionReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request error",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	question := &model.Question{}
	err = db.Mysql.Model(&model.Question{}).Where("id = ?", id).First(&question).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	chapter := &model.Chapter{}
	err = db.Mysql.Model(&model.Chapter{}).Where("id = ?", question.ChapterID).First(&chapter).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	answerChecker, err := answers.GetAnswerChecker(question.Type, question.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "check generate error",
		})
		return
	}

	err = answerChecker.Check(req.Content)
	if err != nil {
		if err == answers.ErrAnswerNotMatch {
			err = db.Mysql.Model(&question).Where("id = ?", question.ID).Update("status", model.FailQuestion).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "update error",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "answer error",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "check error",
		})
		return
	}

	err = db.Mysql.Model(&question).Where("id = ?", question.ID).Update("status", model.PassQuestion).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "update error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
