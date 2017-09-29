package entity

import "database/sql"

//Tag 标签
type Tag struct {
	ID        int64
	TagName   string
	IsDelete  int8   `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

//Scan 读取
func (t *Tag) Scan(row *sql.Row) error {
	return row.Scan(&t.ID, &t.TagName, &t.IsDelete, &t.CreatedAt, &t.UpdatedAt)
}

//ScanRows 读取
func (t *Tag) ScanRows(rows *sql.Rows) error {
	return rows.Scan(&t.ID, &t.TagName, &t.IsDelete, &t.CreatedAt, &t.UpdatedAt)
}

// TagQueryCondition 用户查询条件
type TagQueryCondition struct {
	TagID   sql.NullString
	TagName sql.NullString
	Page
}
