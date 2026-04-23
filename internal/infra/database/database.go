package database

import (
	"database/sql"
	"fmt"
	"go-clean-architecture/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(dbCfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=1&loc=Asia/Seoul",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)

	db, err := open(dbCfg.Type, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbCfg.MaxOpenConns)
	db.SetMaxIdleConns(dbCfg.MaxIdleConns)
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
