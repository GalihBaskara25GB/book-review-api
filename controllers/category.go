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

func GetCategories(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetCategories(database.DbConnection)

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

func GetCategory(c *gin.Context) {
	var (
		result gin.H
	)

	var category structs.Category

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}
	category.Id = int64(id)

	categoryRowData, err := repository.GetCategory(database.DbConnection, category)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categoryRowData,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateCategoryRequest(&category)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateCategoryRequest(&category)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	category.Id = int64(id)
	rowsAffected, err := repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success update category",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Category with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.Id = int64(id)
	rowsAffected, err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success delete category",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Category with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func validateCategoryRequest(category *structs.Category) (errs []string) {
	if category.Name == "" {
		errs = append(errs, "Category name cannot be empty")
	}

	timestamp := time.Now()
	category.UpdatedAt = timestamp
	if category.Id == 0 {
		category.CreatedAt = timestamp
	}

	return
}
