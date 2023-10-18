package operations

import "database/sql"

func CreateWallet(user_id int64, tx *sql.Tx) int64 {
	query := "INSERT INTO wallet(user_id) VALUES(?)"
	res, err := tx.Exec(query, user_id)
	if err != nil {
		return -2
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -2
	}
	return id
}
