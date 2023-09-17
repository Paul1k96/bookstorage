package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Paul1k96/bookstorage/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// NewSqlDB подготовка и выдача драйвера для связи с базой данных MySQL
func NewSqlDB(dbConf config.DB) (*sqlx.DB, error) {
	var err error
	var dsn string
	var dbRaw *sql.DB
	cfg := mysql.NewConfig()
	cfg.Net = dbConf.Net
	cfg.Addr = dbConf.Host
	cfg.User = dbConf.User
	cfg.Passwd = dbConf.Password
	cfg.DBName = dbConf.Name
	cfg.ParseTime = true
	cfg.Timeout = time.Duration(dbConf.Timeout) * time.Second
	dsn = cfg.FormatDSN()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(time.Second * time.Duration(dbConf.Timeout))

	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %d timeout %s", dbConf.Timeout, err)
		case <-ticker.C:
			dbRaw, err = sql.Open(dbConf.Driver, dsn)
			if err != nil {
				return nil, err
			}

			err = dbRaw.Ping()
			if err == nil {
				db := sqlx.NewDb(dbRaw, dbConf.Driver)
				db.SetMaxOpenConns(50)
				db.SetMaxIdleConns(50)
				return db, nil
			}
			log.Println("failed to connect to the database", dsn, err)
		}
	}
}
