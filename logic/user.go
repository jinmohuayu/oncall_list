package logic

import "git.elenet.me/DA/oncall_list/entity"

// GetUser 获取指定ID的用户
func (s Service) GetUser(id int64) (*entity.User, error) {
	return s.userRepository.GetByID(id)
}

// QueryUser 查询用户
func (s Service) QueryUser(condition *entity.UserQueryCondition) error {
	return s.userRepository.Query(condition)
}
