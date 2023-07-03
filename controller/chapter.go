package controller

import (
	"YangCodeCamp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllChaptersResp struct {
	Chapters  []*model.Chapter `json:"chapters"`
	TotalPage int              `json:"total_page"`
	PageNum   int              `json:"page_num"`
}

func GetAllChapters(c *gin.Context) {
	resp := &GetAllChaptersResp{
		Chapters: make([]*model.Chapter, 0),
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}

type GetChapterByIdResp struct {
	Chapter          *model.Chapter `json:"chapter"`
	TotalQuestion    int            `json:"total_question"`
	FinishedQuestion int            `json:"finished_question"`
}

func GetChapterById(c *gin.Context) {

	//idStr := c.Param("id")
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "id error",
	//	})
	//	return
	//}
	//
	//chapter := &model.Chapter{}
	//err = db.Mysql.Model(&model.Chapter{}).Where("id = ?", id).Preload(clause.Associations).First(&chapter).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{
	//			"message": "not found",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "error",
	//	})
	//	return
	//}

	resp := &GetChapterByIdResp{
		Chapter: &model.Chapter{},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}

type GetQuestionByChapterIdResp struct {
	Questions []*model.Question `json:"questions"`
	TotalPage int               `json:"total_page"`
	PageNum   int               `json:"page_num"`
}

func GetQuestionByChapterId(c *gin.Context) {

	//idStr, _ := c.Params.Get("id")
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "id error",
	//	})
	//	return
	//}
	//
	//chapter := &model.Chapter{}
	//err = db.Mysql.Model(&model.Chapter{}).Where("id = ?", id).First(&chapter).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{
	//			"message": "not found",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "error",
	//	})
	//	return
	//}
	//
	//question := &model.Question{}
	//err = db.Mysql.Model(&model.Question{}).Where("id = ?", chapter.QuestionID).First(&question).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{
	//			"message": "not found",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"message": "error",
	//	})
	//	return
	//}

	resp := &GetQuestionByChapterIdResp{
		Questions: make([]*model.Question, 0),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resp,
	})
}
