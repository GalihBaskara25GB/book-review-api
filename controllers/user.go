package controllers

import (
	"bookreview/database"
	"bookreview/repository"
	"bookreview/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetUsers(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	var (
		result gin.H
	)

	var user structs.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}
	user.Id = int64(id)

	userRowData, err := repository.GetUser(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": userRowData,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
			"err":    err,
		})
		panic(err)
	}

	requestValirationError := validateUserRequest(&user, "POST")
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert user",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Username cannot be updated, please remove it from request body",
		})
		return
	}

	requestValirationError := validateUserRequest(&user, "PUT")
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	user.Id = int64(id)
	rowsAffected, err := repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success update user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "User with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.Id = int64(id)
	rowsAffected, err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success delete user",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "User with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func validateUserRequest(user *structs.User, method string) (errs []string) {
	if method == "POST" {
		if user.Username == "" {
			errs = append(errs, "Username cannot be empty")
		} else {
			userRowData, err := repository.GetUserByUsername(database.DbConnection, user.Username)
			if err != nil {
				errs = append(errs, err.Error())
			}

			if len(userRowData) > 0 {
				errs = append(errs, "Username already taken, please choose other username")
			}
		}
	}

	if user.Role == "" {
		errs = append(errs, "Role cannot be empty")
	} else {
		if user.Role != "superuser" && user.Role != "author" && user.Role != "reviewer" {
			errs = append(errs, "Invalid user role, user role must be either 'superuser', 'reviewer' or 'author'")
		}
	}

	timestamp := time.Now()
	user.UpdatedAt = timestamp
	if user.Id == 0 {
		user.CreatedAt = timestamp
	}

	return
}
