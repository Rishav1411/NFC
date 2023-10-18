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

func CheckBalance(wallet_id int64, db *sql.DB) int64 {
	query := "SELECT balance FROM wallet WHERE wallet_id = ?"
	var val int64
	err := db.QueryRow(query, wallet_id).Scan(&val)
	if err != nil {
		return -1
	}
	return val
}

func UpdateBalance(wallet_id int64, balance int64, tx *sql.Tx) error {
	query := "UPDATE wallet SET balance = balance + ? WHERE wallet_id = ?"
	_, err := tx.Exec(query, balance, wallet_id)
	return err
}
