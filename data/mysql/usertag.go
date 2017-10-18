package mysql

import "database/sql"

// UserTagRepository 用户Repository
type UserTagRepository struct {
	db *sql.DB
}

// NewUserTagRepository 新建用户Repository
func NewUserTagRepository(db *sql.DB) *UserTagRepository {
	return &UserTagRepository{db: db}
}

// DeleteByUserID 通过指定用户的ID来删除关联表user_tag的整条记录
func (s UserTagRepository) DeleteByUserID(tx *sql.Tx, userID int64) error {
	_, err := tx.Exec("delete from user_tag where user_id = ?", userID)

	return err
}

//Insert 删除之后再插入user_tag的所有信息
func (s UserTagRepository) Insert(tx *sql.Tx, userID, tagID int64) error {
	_, err := tx.Exec("insert into user_tag(user_id,tag_id) values (?, ?)", userID, tagID)

	return err
}
