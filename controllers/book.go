package controllers

import (
	"bookreview/database"
	"bookreview/repository"
	"bookreview/structs"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetBook(c *gin.Context) {
	var (
		result gin.H
	)

	var book structs.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}
	book.Id = int64(id)

	bookRowData, err := repository.GetBook(database.DbConnection, book)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": bookRowData,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateBookRequest(&book)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		fmt.Print(requestValirationError[0])
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateBookRequest(&book)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	book.Id = int64(id)
	rowsAffected, err := repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success update book",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Book with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.Id = int64(id)
	rowsAffected, err := repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success delete book",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Book with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func validateBookRequest(book *structs.Book) (errs []string) {
	if book.CategoryId != 0 {
		var category structs.Category

		categoryId := strconv.Itoa(book.CategoryId)
		category.Id = int64(book.CategoryId)

		categoryRowData, err := repository.GetCategory(database.DbConnection, category)
		if err != nil {
			errs = append(errs, err.Error())
		}

		if len(categoryRowData) <= 0 {
			errs = append(errs, "Category with ID:"+categoryId+" IS NOT FOUND, please choose other category")
		}
	} else {
		errs = append(errs, "Category cannot be empty")
	}

	if book.ReleaseYear < 0 {
		errs = append(errs, "Release year cannot be empty")
	}

	_, err := url.ParseRequestURI(book.ImageUrl)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if book.Author == "" {
		errs = append(errs, "Book author cannot be empty")
	}
	if book.Title == "" {
		errs = append(errs, "Book title cannot be empty")
	}
	if book.TotalPage <= 0 {
		errs = append(errs, "Total page cannot be empty")
	}

	timestamp := time.Now()
	book.UpdatedAt = timestamp
	if book.Id == 0 {
		book.CreatedAt = timestamp
	}

	return
}
