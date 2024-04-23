package repository

import (
	"bookreview/structs"
	"database/sql"
)

const tableCategory = "categories"

func GetCategories(db *sql.DB) (results []structs.Category, err error) {
	sql := "SELECT * FROM " + tableCategory

	rows, err := db.Query(sql)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var category = structs.Category{}
		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, category)
	}

	return
}

func GetCategory(db *sql.DB, category structs.Category) (results []structs.Category, err error) {
	sql := "SELECT * FROM " + tableCategory + " WHERE id = $1"

	rows, err := db.Query(sql, category.Id)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var category = structs.Category{}
		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `
		INSERT INTO ` + tableCategory + ` (name, created_at, updated_at)
		VALUES ($1, $2, $3)
	`

	_, err = db.Exec(sql, category.Name, category.CreatedAt, category.UpdatedAt)

	return
}

func UpdateCategory(db *sql.DB, category structs.Category) (rowsAffected int64, err error) {
	sql := `
		UPDATE ` + tableCategory + `
		SET name = $2, updated_at = $3
		WHERE id = $1
	`
	res, err := db.Exec(sql, category.Id, category.Name, category.UpdatedAt)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteCategory(db *sql.DB, category structs.Category) (rowsAffected int64, err error) {
	sql := `
		DELETE FROM ` + tableCategory + `
		WHERE id = $1
	`

	res, err := db.Exec(sql, category.Id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return
}
