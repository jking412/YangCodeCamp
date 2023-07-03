package web

import (
	"YangCodeCamp/controller"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	r.GET("/ping", ping)

	r.GET("/course", controller.GetAllCourses)
	courseGroup := r.Group("/course")
	{
		courseGroup.GET("/:id", controller.GetCourseById)
		courseGroup.GET("/:id/chapter", controller.GetChaptersByCourseId)
		//courseGroup.GET("/:id/progress", controller.GetProgressByCourseId)
		//courseGroup.GET("/:id/reset", controller.ResetChapters)
		//courseGroup.GET("/:id/resume", controller.ResumeChapter)
	}

	r.GET("/chapter", controller.GetAllChapters)
	chapterGroup := r.Group("/chapter")
	{
		chapterGroup.GET("/:id", controller.GetChapterById)
		chapterGroup.GET("/:id/question", controller.GetQuestionByChapterId)
	}

	questionGroup := r.Group("/question")
	{
		questionGroup.GET("/:id", controller.GetQuestionById)
		questionGroup.POST("/:id/submit", controller.SubmitQuestion)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
