package entity

import "database/sql"

// User 用户
type User struct {
	ID   int64
	Name string
}

// Scan 读取
func (u *User) Scan(row *sql.Row) error {
	return row.Scan(&u.ID, &u.Name)
}

// ScanRows 读取
func (u *User) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&u.ID, &u.Name)
}

// UserQueryCondition 用户查询条件
type UserQueryCondition struct {
	Name sql.NullString
	Page
}
