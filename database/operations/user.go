package operations

import (
	"database/sql"
)

func CheckUser(phone string, db *sql.DB) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE phone_number = ?"
	err := db.QueryRow(query, phone).Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return -1, nil
		}
		return -2, err
	}
	return id, nil
}

func RegisterUser(phone string, user_name string, reg string, db *sql.DB) int {
	query := "INSERT INTO users(user_name,reg_number,phone_number) VALUES(?,?,?)"
	err := db.QueryRow(query, user_name, reg, phone).Err()
	if err != nil {
		return -1
	}
	var id int
	id, err = CheckUser(phone, db)
	if err != nil {
		return -1
	}
	return id
}
