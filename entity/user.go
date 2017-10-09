package entity

import (
	"database/sql"
	"time"
)

// User 用户
type User struct {
	ID         int64
	Name       string
	Department string
	Product    string
	Email      string
	PhoneNum   string
	Remark     string
	IsDelete   int8      `json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Tags       []Tag     `json:"tags"` //用户标签
}

// Scan 读取
func (u *User) Scan(row *sql.Row) error {
	return row.Scan(&u.ID, &u.Name, &u.Department, &u.Product, &u.Email, &u.PhoneNum, &u.Remark, &u.IsDelete, &u.CreatedAt, &u.UpdatedAt)
}

// ScanRows 读取
func (u *User) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&u.ID, &u.Name, &u.Department, &u.Product, &u.Email, &u.PhoneNum, &u.Remark, &u.IsDelete, &u.CreatedAt, &u.UpdatedAt)
}

// UserQueryCondition 用户查询条件
type UserQueryCondition struct {
	Name sql.NullString
	Page
}

// UserPrivileges 用户特权
type UserPrivileges int

const (
	// AllowEdit 允许编辑
	AllowEdit UserPrivileges = iota + 1
)

// UserDetail 当前用户完整信息
type UserDetail struct {
	*User
	Privileges UserPrivileges
}
