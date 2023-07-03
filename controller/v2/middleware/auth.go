package middleware

import (
	"YangCodeCamp/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Jwt解析失败",
			})
			c.Abort()
		}
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Next()
	}
}
