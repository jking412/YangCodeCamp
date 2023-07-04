package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/answers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	err = db.Mysql.Model(&model.Question{}).Where("id = ?", id).
		Preload(clause.Associations).First(&question).Error
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
	Code []struct {
		Type    model.Status `json:"type"`
		Content string       `json:"content"`
	} `json:"code"`
}

type SubmitQuestionResp struct {
	Result model.Status `json:"result"`
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
	err = db.Mysql.Model(&model.Question{}).Where("id = ?", id).
		Preload(clause.Associations).First(&question).Error
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

	// check answer
	for _, questionDetail := range question.Detail {
		var answerChecker answers.Answer
		answerChecker, err = answers.GetAnswerChecker(questionDetail.Type, questionDetail.Answer, questionDetail.CheckMessage)
		if err != nil {
			goto outCheck
		}
		for _, code := range req.Code {
			if int(code.Type) == int(questionDetail.Type) {
				err = answerChecker.Check(code.Content)
				if err != nil {
					goto outCheck
				}
			}
		}
	}

	// goto
outCheck:
	;

	// 根据err来判断是否正确
	if err != nil {
		if err == answers.ErrAnswerNotMatch {
			err = db.Mysql.Model(&question).Where("id = ?", question.ID).Update("status", model.FailQuestion).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "update error",
				})
				return
			}
			c.JSON(http.StatusOK, SubmitQuestionResp{
				Result: model.FailQuestion,
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
	c.JSON(http.StatusOK, SubmitQuestionResp{
		Result: model.PassQuestion,
	})
}
