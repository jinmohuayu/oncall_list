package logic

import "git.elenet.me/DA/oncall_list/entity"

// QueryTag 查询标签
func (s Service) QueryTag(condition *entity.TagQueryCondition) error {
	return s.tagRepository.TagQuery(condition)
}
