package controllers

import (
	"bookreview/database"
	"bookreview/repository"
	"bookreview/structs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	var (
		result gin.H
	)

	reviews, err := repository.GetReviews(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": reviews,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetReview(c *gin.Context) {
	var (
		result gin.H
	)

	var review structs.Review

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}
	review.Id = int64(id)

	reviewRowData, err := repository.GetReview(database.DbConnection, review)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": reviewRowData,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertReview(c *gin.Context) {
	var review structs.Review

	err := c.ShouldBindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateReviewRequest(&review)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		fmt.Print(requestValirationError[0])
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	err = repository.InsertReview(database.DbConnection, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert review",
	})
}

func UpdateReview(c *gin.Context) {
	var review structs.Review

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	requestValirationError := validateReviewRequest(&review)
	if requestValirationError != nil || len(requestValirationError) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": requestValirationError,
		})
		return
	}

	review.Id = int64(id)
	rowsAffected, err := repository.UpdateReview(database.DbConnection, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success update review",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Review with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func DeleteReview(c *gin.Context) {
	var review structs.Review
	id, _ := strconv.Atoi(c.Param("id"))

	review.Id = int64(id)
	rowsAffected, err := repository.DeleteReview(database.DbConnection, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": "Opps something went wrong",
		})
		panic(err)
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "Success delete review",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Review with id:" + strconv.Itoa(id) + " IS NOT FOUND",
		})
	}
}

func validateReviewRequest(review *structs.Review) (errs []string) {
	if review.UserId != 0 {
		var user structs.User

		user.Id = int64(review.UserId)

		userRowData, err := repository.GetUser(database.DbConnection, user)
		if err != nil {
			errs = append(errs, err.Error())
		}

		if len(userRowData) <= 0 {
			errs = append(errs, "User with ID:"+strconv.Itoa(int(review.UserId))+" IS NOT FOUND, please choose other user")
		}
	} else {
		errs = append(errs, "User cannot be empty")
	}

	if review.BookId != 0 {
		var book structs.Book

		book.Id = int64(review.BookId)

		bookRowData, err := repository.GetBook(database.DbConnection, book)
		if err != nil {
			errs = append(errs, err.Error())
		}

		if len(bookRowData) <= 0 {
			errs = append(errs, "Book with ID:"+strconv.Itoa(int(review.BookId))+" IS NOT FOUND, please choose other book")
		}
	} else {
		errs = append(errs, "Book cannot be empty")
	}

	if review.Stars < 0 || review.Stars > 5 {
		errs = append(errs, "Stars must be between 0 and 5")
	}

	timestamp := time.Now()
	review.UpdatedAt = timestamp
	if review.Id == 0 {
		review.CreatedAt = timestamp
	}

	return
}
