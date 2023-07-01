package web

import "github.com/gin-gonic/gin"

func initRouter() {
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
