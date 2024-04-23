package repository

import (
	"bookreview/structs"
	"database/sql"
)

const tableBook = "books"

func GetBooks(db *sql.DB) (results []structs.Book, err error) {
	sql := "SELECT * FROM " + tableBook

	rows, err := db.Query(sql)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var book = structs.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Author, &book.CreatedAt, &book.UpdatedAt, &book.CategoryId)
		if err != nil {
			break
		}
		results = append(results, book)
	}

	return
}

func GetBook(db *sql.DB, book structs.Book) (results []structs.Book, err error) {
	sql := "SELECT * FROM " + tableBook + " WHERE id = $1"

	rows, err := db.Query(sql, book.Id)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var book = structs.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.ImageUrl, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Author, &book.CreatedAt, &book.UpdatedAt, &book.CategoryId)
		if err != nil {
			break
		}
		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := `
		INSERT INTO ` + tableBook + ` (title, description, image_url, release_year, price, total_page, category_id, author, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(sql, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.CategoryId, book.Author, book.CreatedAt, book.UpdatedAt)

	return
}

func UpdateBook(db *sql.DB, book structs.Book) (rowsAffected int64, err error) {
	sql := `
		UPDATE ` + tableBook + `
		SET title = $2,
		description = $3,
		image_url = $4,
		release_year = $5,
		price = $6,
		total_page = $7,
		category_id = $8,
		author = $9,
		updated_at = $10
		WHERE id = $1
	`
	res, err := db.Exec(sql, book.Id, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.CategoryId, book.Author, book.UpdatedAt)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteBook(db *sql.DB, book structs.Book) (rowsAffected int64, err error) {
	sql := `
		DELETE FROM ` + tableBook + `
		WHERE id = $1
	`

	res, err := db.Exec(sql, book.Id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}
