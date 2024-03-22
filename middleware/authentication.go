package middleware

import (
	"agit-test/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "403",
				"status":  "Forbidden",
				"message": "Unauthorized",
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
