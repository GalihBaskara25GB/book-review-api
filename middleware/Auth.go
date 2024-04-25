package middleware

import (
	"bookreview/database"
	"bookreview/repository"
	"bookreview/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(allowedRole []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		resultString := ""
		responseHeader := http.StatusBadRequest

		username, password, ok := c.Request.BasicAuth()

		if !ok {
			resultString = "username and password is required"
		}

		if ok {
			var user structs.User
			user.Username = username
			user.Password = password

			userRowData, err := repository.AuthenticateUser(database.DbConnection, user)

			if err != nil {
				resultString = err.Error()
			} else {
				if len(userRowData) > 0 {
					isAuthorized := false

					for _, role := range allowedRole {
						if role == "superuser" || role == "reviewer" || role == "author" {
							if role == userRowData[0].Role {
								isAuthorized = true
							}
						}
					}

					if isAuthorized || len(allowedRole) == 0 {
						c.Next()
						return
					} else {
						resultString = "you're not allowed to access this route"
						responseHeader = http.StatusUnauthorized
					}

				} else {
					resultString = "username and password is incorrect"
				}
			}
		}

		c.JSON(responseHeader, gin.H{
			"result": resultString,
		})
		c.Abort()
	}
}
