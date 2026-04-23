package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	MaxOpenConns int
	MaxIdleConns int
}

func New(driver, dsn string, opts Options) (*sql.DB, error) {
	db, err := open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(opts.MaxOpenConns)
	db.SetMaxIdleConns(opts.MaxIdleConns)
	return db, nil
}

func open(dbms string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbms, dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 3; i++ {
		if err := db.Ping(); err == nil {
			return db, nil
		}
		time.Sleep(1 * time.Second)
	}
	return nil, fmt.Errorf("failed to ping database after retries")
}
