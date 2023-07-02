package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetChapterByClassId(c *gin.Context) {

	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "id error",
		})
		return
	}

	chapters := make([]*model.Chapter, 0)
	err = db.Mysql.Model(&model.Chapter{}).Where("class_id = ?", id).Find(&chapters).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    chapters,
	})
}
