package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func SQLConnection() *sql.DB {
	maxRetries := 3
	sleepDuration := time.Second * 5
	for retries := 0; retries < maxRetries; retries++ {
		db, err := sql.Open("mysql", SQL_STRING)
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

func RedisConnection() *redis.Client {
	maxRetries := 3
	for retries := 0; retries < maxRetries; retries++ {
		opt, err := redis.ParseURL(REDIS_STRING)
		if err == nil {
			return redis.NewClient(opt)
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}
