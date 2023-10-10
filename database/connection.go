package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	maxRetries := 3
	sleepDuration := time.Second * 5
	for retries := 0; retries < maxRetries; retries++ {
		db, err := sql.Open("mysql", CONNECTION_STRING)
		if err != nil {
			time.Sleep(sleepDuration)
			continue
		}
		err = db.Ping()
		if err != nil {
			time.Sleep(sleepDuration)
			continue
		}
		return db
	}
	return nil
}
