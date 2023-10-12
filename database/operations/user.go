package operations

import "database/sql"

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
