package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func GetAllCourses(c *gin.Context) {
	courses := make([]*model.Course, 0)

	err := db.Mysql.Model(&model.Course{}).Find(&courses).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    courses,
	})
}

func GetCourseById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	course := &model.Course{}
	err = db.Mysql.Model(&model.Course{}).Where("id = ?", id).Preload(clause.Associations).First(&course).Error
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    course,
	})
}

func GetChaptersByCourseId(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	chapters := make([]*model.Chapter, 0)
	err = db.Mysql.Model(&model.Chapter{}).Where("course_id = ?", id).Find(&chapters).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    chapters,
	})
}
