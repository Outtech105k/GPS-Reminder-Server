package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	conn := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Addr:      "mysql-db",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, err := sql.Open("mysql", conn.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
