package repository

import (
	"bookreview/structs"
	"database/sql"
)

const tableUser = "users"

func GetUsers(db *sql.DB) (results []structs.User, err error) {
	sql := "SELECT * FROM " + tableUser

	rows, err := db.Query(sql)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, user)
	}

	return
}

func GetUser(db *sql.DB, user structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM " + tableUser + " WHERE id = $1"

	rows, err := db.Query(sql, user.Id)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, user)
	}

	return
}

func GetUserByUsername(db *sql.DB, username string) (results []structs.User, err error) {
	sql := "SELECT * FROM " + tableUser + " WHERE username = '$1'"

	rows, err := db.Query(sql, username)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, user)
	}

	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := `
		INSERT INTO ` + tableUser + ` (username, password, role, created_at, updated_at)
		VALUES ($1, md5($2), $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(sql, user.Username, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	return
}

func UpdateUser(db *sql.DB, user structs.User) (rowsAffected int64, err error) {
	sql := `
		UPDATE ` + tableUser + `
		SET password = $2,
		role = $3,
		updated_at = $4
		WHERE id = $1
	`
	res, err := db.Exec(sql, user.Id, user.Password, user.Role, user.UpdatedAt)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteUser(db *sql.DB, user structs.User) (rowsAffected int64, err error) {
	sql := `
		DELETE FROM ` + tableUser + `
		WHERE id = $1
	`

	res, err := db.Exec(sql, user.Id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

func AuthenticateUser(db *sql.DB, user structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM " + tableUser + " WHERE username = $1 AND password = md5($2)"

	rows, err := db.Query(sql, user.Username, user.Password)

	if err != nil {
		defer rows.Close()
	}

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			break
		}
		results = append(results, user)
	}

	return
}
