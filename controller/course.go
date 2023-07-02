package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"github.com/gin-gonic/gin"
)

func GetAllClasses(c *gin.Context) {
	classes := make([]*model.Class, 0)
	err := db.Mysql.Model(&model.Class{}).Find(&classes).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    classes,
	})
}
