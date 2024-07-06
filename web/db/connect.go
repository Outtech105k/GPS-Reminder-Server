package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tanimutomo/sqlfile"
)

func Connect() (*sql.DB, error) {
	// データベース接続
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("timeset: %w", err)
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
		return nil, fmt.Errorf("dbConnection: %w", err)
	}

	// テーブル構成初期化
	s := sqlfile.New()

	err = s.File("db/init.sql")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("readInitDbFile: %w", err)
	}
	_, err = s.Exec(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("execInitDbFile: %w", err)
	}

	return db, nil
}
