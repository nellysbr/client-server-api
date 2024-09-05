package database

import (
	"context"
	"database/sql"
	"time"
)

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS quotations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid REAL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	return err
}

func SaveQuotation(ctx context.Context, db *sql.DB, bid float64) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO quotations (bid) VALUES (?)", bid)
	return err
}