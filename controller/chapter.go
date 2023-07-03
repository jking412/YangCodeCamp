package controller

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/paginations"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type GetAllChaptersResp struct {
	Chapters  []*model.Chapter `json:"chapters"`
	TotalPage int              `json:"total_page"`
	PageNum   int              `json:"page_num"`
}

func GetAllChapters(c *gin.Context) {

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

	chapters := make([]*model.Chapter, 0)
	err = db.Mysql.Model(&model.Chapter{}).Find(&chapters).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	pagination := paginations.NewPagination(len(chapters), pageNum, pageSize)
	chapters = chapters[pagination.Start():pagination.End()]

	resp := &GetAllChaptersResp{
		Chapters:  chapters,
		TotalPage: pagination.TotalPage(),
		PageNum:   pagination.PageNum,
	}
	c.JSON(http.StatusOK, resp)
}

type GetChapterByIdResp struct {
	Chapter          *model.Chapter `json:"chapter"`
	TotalQuestion    int            `json:"total_question"`
	FinishedQuestion int            `json:"finished_question"`
}

func GetChapterById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

	chapter := &model.Chapter{}
	err = db.Mysql.Model(&model.Chapter{}).Where("id = ?", id).First(&chapter).Error
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

	resp := &GetChapterByIdResp{
		Chapter: chapter,
	}

	c.JSON(http.StatusOK, resp)
}

type GetQuestionByChapterIdResp struct {
	Questions []*model.Question `json:"questions"`
	TotalPage int               `json:"total_page"`
	PageNum   int               `json:"page_num"`
}

func GetQuestionByChapterId(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id error",
		})
		return
	}

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

	questions := make([]*model.Question, 0)
	err = db.Mysql.Model(&model.Question{}).Where("chapter_id = ?", id).Find(&questions).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	pagination := paginations.NewPagination(len(questions), pageNum, pageSize)
	questions = questions[pagination.Start():pagination.End()]

	resp := &GetQuestionByChapterIdResp{
		Questions: questions,
		TotalPage: pagination.TotalPage(),
		PageNum:   pagination.PageNum,
	}

	c.JSON(http.StatusOK, resp)
}
