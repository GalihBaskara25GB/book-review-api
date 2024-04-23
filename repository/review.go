package repository

import (
	"bookreview/structs"
	"database/sql"
)

const tableReview = "reviews"

func GetReviews(db *sql.DB) (results []structs.Review, err error) {
	sql := "SELECT * FROM " + tableReview

	rows, err := db.Query(sql)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var review = structs.Review{}
		err = rows.Scan(&review.Id, &review.UserId, &review.BookId, &review.Description, &review.Stars, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, review)
	}

	return
}

func GetReview(db *sql.DB, review structs.Review) (results []structs.Review, err error) {
	sql := "SELECT * FROM " + tableReview + " WHERE id = $1"

	rows, err := db.Query(sql, review.Id)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var review = structs.Review{}
		err = rows.Scan(&review.Id, &review.UserId, &review.BookId, &review.Description, &review.Stars, &review.CreatedAt, &review.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, review)
	}

	return
}

func InsertReview(db *sql.DB, review structs.Review) (err error) {
	sql := `
		INSERT INTO ` + tableReview + ` (user_id, book_id, description, stars, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(sql, review.UserId, review.BookId, review.Description, review.Stars, review.CreatedAt, review.UpdatedAt)

	return
}

func UpdateReview(db *sql.DB, review structs.Review) (rowsAffected int64, err error) {
	sql := `
		UPDATE ` + tableReview + `
		SET user_id = $2,
		book_id = $3,
		description = $4,
		stars = $5,
		updated_at = $6
		WHERE id = $1
	`
	res, err := db.Exec(sql, review.Id, review.UserId, review.BookId, review.Description, review.Stars, review.UpdatedAt)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteReview(db *sql.DB, review structs.Review) (rowsAffected int64, err error) {
	sql := `
		DELETE FROM ` + tableReview + `
		WHERE id = $1
	`

	res, err := db.Exec(sql, review.Id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}
