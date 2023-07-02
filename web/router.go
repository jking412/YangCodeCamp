package web

import (
	"YangCodeCamp/controller"
	"github.com/gin-gonic/gin"
)

func initRouter() {
	r.GET("/ping", ping)

	r.GET("/class", controller.GetAllClasses)
	//classGroup := r.Group("/class")
	//{
	//}

	chapterGroup := r.Group("/chapter")
	{
		chapterGroup.GET("/:id", controller.GetChapterByClassId)
	}

	questionGroup := r.Group("/question")
	{
		questionGroup.GET("/:id", controller.GetQuestionById)
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
