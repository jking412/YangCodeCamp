package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/paginations"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

type GetAllCoursesResp struct {
	Courses   []*model.Course `json:"courses"`
	TotalPage int             `json:"total_page"`
	PageNum   int             `json:"page_num"`
}

func GetAllCourses(c *gin.Context) {

	pageNumStr := c.Query("page_num")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		pageNum = 1
	}
	pageSizeStr := c.Query("page_size")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	var count int64
	err = db.Mysql.Model(&model.Course{}).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	courses := make([]*model.Course, 0)
	pagination := paginations.NewPagination(int(count), pageNum, pageSize)
	err = db.Mysql.Model(&model.Course{}).Offset(pagination.Offset()).Limit(pagination.Limit()).Find(&courses).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	resp := &GetAllCoursesResp{
		Courses:   courses,
		TotalPage: pagination.TotalPage(),
		PageNum:   pagination.PageNum,
	}

	c.JSON(http.StatusOK, resp)
}

type GetCourseByIdResp struct {
	Course           *model.Course `json:"course"`
	TotalQuestion    int           `json:"total_question"`
	FinishedQuestion int           `json:"finished_question"`
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

	resp := &GetCourseByIdResp{
		Course: &model.Course{},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}

type GetChaptersByCourseIdResp struct {
	Chapters  []*model.Chapter `json:"chapters"`
	TotalPage int              `json:"total_page"`
	PageNum   int              `json:"page_num"`
}

func GetChaptersByCourseId(c *gin.Context) {

	//idStr := c.Param("id")
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "id error",
	//	})
	//	return
	//}
	//
	//chapters := make([]*model.Chapter, 0)
	//err = db.Mysql.Model(&model.Chapter{}).Where("course_id = ?", id).Find(&chapters).Error
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "error",
	//	})
	//	return
	//}
	resp := &GetChaptersByCourseIdResp{
		Chapters: make([]*model.Chapter, 0),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}

func GetProgressByCourseId(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	chapter := make([]*model.Chapter, 0)
	err = db.Mysql.Model(&model.Chapter{}).Where("course_id = ?", id).Find(&chapter).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	sum := len(chapter)
	completeCount := 0
	for _, v := range chapter {
		if v.Status == model.SuccessChapter {
			completeCount++
		}
	}

	completeRate := float64(completeCount) / float64(sum)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"complete-rate": completeRate,
		},
	})

}

//func ResetChapters(c *gin.Context) {
//
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "id error",
//		})
//		return
//	}
//
//	err = db.Mysql.Model(&model.Chapter{}).Where("course_id = ?", id).Update("status", model.UnfinishedChapter).Error
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"message": "error",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "success",
//	})
//
//}

// ResumeChapter 获取最前面的未完成章节
//func ResumeChapter(c *gin.Context) {
//
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "id error",
//		})
//		return
//	}
//
//	chapter := &model.Chapter{}
//	err = db.Mysql.Model(&model.Chapter{}).Where("course_id = ? AND status = ?", id, model.UnfinishedChapter).
//		Order("number asc").First(&chapter).Error
//
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//			c.JSON(http.StatusNotFound, gin.H{
//				"message": "not found",
//			})
//			return
//		}
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"message": "error",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "success",
//		"data":    chapter,
//	})
//}
