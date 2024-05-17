// Package mariadb по сути просто бойлерплейта по созданию инстанса клиента БД
package mariadb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	maxConnectionLifeTime = time.Minute * 3
	maxOpenConns          = 100
	maxIdleConns          = 100
)

// New создает стандартный инстанс клиента БД
func New(host, port, username, password, database string) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", host, port),
		DBName:               database,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sql.Open failed: %w", err)
	}

	db.SetConnMaxLifetime(maxConnectionLifeTime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("sql.Ping failed: %w", pingErr)
	}

	return db, nil
}
