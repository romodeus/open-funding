package utils

import (
	"database/sql"
	"fmt"
	"log"
	"open-funding/config"

	_ "github.com/lib/pq"
)

func InitDB(cfg *config.AppConfig) *sql.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=verify-full", cfg.DB_NAME, cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_USER)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
