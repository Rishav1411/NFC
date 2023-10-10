package operations

import "database/sql"

func CheckUser(phone string, db *sql.DB) (bool, error) {
	var id int
	query := "SELECT COUNT(*) FROM users WHERE phone_number = ?"
	err := db.QueryRow(query, phone).Scan(&id)
	if err != nil {
		return false, err
	}
	return id > 0, nil
}
