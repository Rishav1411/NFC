package operations

import (
	"database/sql"
)

func CheckUser(phone string, db *sql.DB) int64 {
	var id int64
	query := "SELECT user_id FROM users WHERE phone_number = ?"
	err := db.QueryRow(query, phone).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return -1
		}
		return -2
	}
	return id
}

func RegisterUser(phone string, user_name string, reg string, tx *sql.Tx) int64 {
	query := "INSERT INTO users(user_name,reg_number,phone_number) VALUES(?,?,?)"
	res, err := tx.Exec(query, user_name, reg, phone)
	if err != nil {
		return -2
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -2
	}
	return id
}

func CheckReg(reg string, db *sql.DB) int64 {
	var id int64
	query := "SELECT user_id FROM users WHERE reg_number = ?"
	err := db.QueryRow(query, reg).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return -1
		}
		return -2
	}
	return id
}
