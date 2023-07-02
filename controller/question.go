package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQuestionById(c *gin.Context) {

	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "id error",
		})
		return
	}

	questions := make([]*model.Question, 0)
	err = db.Mysql.Model(&model.Question{}).Where("chapter_id = ?", id).Find(&questions).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    questions,
	})
}
