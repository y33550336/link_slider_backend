package database

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectToDatabase(mysqlCfg *mysql.Config) (*sqlx.DB, error) {
	if mysqlCfg == nil {
		return nil, errors.New("config is nil")
	}

	db, err := sqlx.Connect("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close() // Close the connection if ping fails
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := migrateTables(db.DB); err != nil {
		_ = db.Close() // Close the connection if migration fails
		return nil, fmt.Errorf("failed to migrate tables: %w", err)
	}

	return db, nil
}
