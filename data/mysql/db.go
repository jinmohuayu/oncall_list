package mysql

import (
	"database/sql"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// GetDB 获取mysql连接
func GetDB(conn string) (*sql.DB, error) {
	return sql.Open("mysql", conn)
}
