package logic

import (
	"database/sql"
	"fmt"
	"git.elenet.me/DA/oncall_list/entity"
)

// GetUser 获取指定ID的用户
func (s Service) GetUser(id int64) (*entity.User, error) {

	// 获取到前端传过来的id
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	//查询用户关联的标签
	tagQueryCondition := entity.TagQueryCondition{UserID: sql.NullInt64{Int64: id, Valid: true}}
	err = s.tagRepository.TagQuery(&tagQueryCondition)
	if err != nil {
		return nil, err
	}

	user.Tags = tagQueryCondition.Records.([]entity.Tag)
	return user, nil
}

// QueryUser 查询用户
func (s Service) QueryUser(condition *entity.UserQueryCondition) error {

	err := s.userRepository.Query(condition)
	if err != nil {
		return err
	}

	users := condition.Records.([]entity.User)
	tagQueryCondition := entity.TagQueryCondition{UserID: sql.NullInt64{Valid: true}}
	for index, user := range users {
		tagQueryCondition.UserID.Int64 = user.ID
		err = s.tagRepository.TagQuery(&tagQueryCondition)
		if err != nil {
			return err
		}

		users[index].Tags = tagQueryCondition.Records.([]entity.Tag)
	}

	condition.Records = users
	return nil
}

// UpdateUserTag 更新用户Tag
func (s Service) UpdateUserTag(userID int64, tagNames []string) error {

	tx, err := s.db.Begin()
	if err != nil {
		//return fmt.Errorf("1: %v", err)
		return err
	}

	err = s.updateUserTagPrivate(tx, userID, tagNames)
	if err != nil {
		s.logger.Printf("updateUserTagPrivate=%v", err)
		err1 := tx.Rollback()
		if err1 != nil {
			return fmt.Errorf("3: %v", err)
		}

		return fmt.Errorf("2: %v", err)
	}

	return tx.Commit()
}

//
func (s Service) updateUserTagPrivate(tx *sql.Tx, userID int64, tagNames []string) error {

	// tagnames数组转变为tagid数组
	tagIDs, err := s.ensureTags(tx, tagNames)
	if err != nil {
		return fmt.Errorf("4: %v", err)
	}

	// delete user tag relation
	err = s.userTagRepository.DeleteByUserID(tx, userID)
	if err != nil {
		return fmt.Errorf("5: %v", err)
	}

	// insert new user_tag
	for _, tagID := range tagIDs {

		err = s.userTagRepository.Insert(tx, userID, tagID)
		if err != nil {
			return fmt.Errorf("6: %v", err)
		}
	}

	// delete unused tag
	return s.tagRepository.DeleteUnusedTag(tx)
}
