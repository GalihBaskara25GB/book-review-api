package repository

import (
	"bookreview/structs"
	"database/sql"
)

const tableReview = "reviews"

func GetReviews(db *sql.DB) (results []structs.ReviewWithUserWithBookWithCategory, err error) {
	sql := `SELECT 
					r.*, 
					u.username AS user_username, 
					b.title AS book_title, 
					b.description AS book_description,
					b.image_url AS book_image_url,
					b.release_year AS book_release_year,
					b.price AS book_price,
					b.total_page AS book_total_page,
					b.author AS book_author,
					b.category_id AS category_id,
					c.name AS category_name
					FROM ` + tableReview + ` AS r 
					INNER JOIN ` + tableUser + ` AS u on u.id = r.user_id 
					INNER JOIN ` + tableBook + ` AS b on b.id = r.book_id 
					INNER JOIN ` + tableCategory + ` AS c on c.id = b.category_id `

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
		defer rows.Close()
	}

	for rows.Next() {
		var review = structs.ReviewWithUserWithBookWithCategory{}
		err = rows.Scan(&review.Id, &review.UserId, &review.BookId, &review.Description, &review.Stars, &review.CreatedAt, &review.UpdatedAt,
			&review.UserUsername,
			&review.BookTitle, &review.BookDescription, &review.BookImageUrl, &review.BookReleaseYear, &review.BookPrice, &review.BookTotalPage, &review.BookAuthor,
			&review.CategoryId, &review.CategoryName)
		if err != nil {
			break
		}
		results = append(results, review)
	}

	return
}

func GetReview(db *sql.DB, review structs.Review) (results []structs.ReviewWithUserWithBookWithCategory, err error) {
	sql := `SELECT 
					r.*, 
					u.username AS user_username, 
					b.title AS book_title, 
					b.description AS book_description,
					b.image_url AS book_image_url,
					b.release_year AS book_release_year,
					b.price AS book_price,
					b.total_page AS book_total_page,
					b.author AS book_author,
					b.category_id AS category_id,
					c.name AS category_name
					FROM ` + tableReview + ` AS r 
					INNER JOIN ` + tableUser + ` AS u on u.id = r.user_id 
					INNER JOIN ` + tableBook + ` AS b on b.id = r.book_id 
					INNER JOIN ` + tableCategory + ` AS c on c.id = b.category_id 
					WHERE r.id = $1`

	rows, err := db.Query(sql, review.Id)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var review = structs.ReviewWithUserWithBookWithCategory{}
		err = rows.Scan(&review.Id, &review.UserId, &review.BookId, &review.Description, &review.Stars, &review.CreatedAt, &review.UpdatedAt,
			&review.UserUsername,
			&review.BookTitle, &review.BookDescription, &review.BookImageUrl, &review.BookReleaseYear, &review.BookPrice, &review.BookTotalPage, &review.BookAuthor,
			&review.CategoryId, &review.CategoryName)
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
		VALUES ($1, $2, $3, $4, $5, $6)
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
