package logic

import (
	"database/sql"
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
