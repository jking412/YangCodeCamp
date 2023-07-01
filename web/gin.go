package web

import (
	"YangCodeCamp/pkg/config"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() {
	r = gin.Default()
	initRouter()
	err := r.Run(":" + config.GetString("web.port"))
	if err != nil {
		panic(err)
	}
}
