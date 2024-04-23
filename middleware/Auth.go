package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		resultString := ""
		username, password, ok := c.Request.BasicAuth()

		if !ok {
			resultString = "username and password is required"
		}

		if (username == "admin" && password == "password") || (username == "editor" && password == "secret") {
			c.Next()
			return
		}
		if ok {
			resultString = "username and password is incorrect"
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"result": resultString,
		})
		c.Abort()
	}
}
