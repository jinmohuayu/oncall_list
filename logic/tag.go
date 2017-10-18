package logic

import (
	"database/sql"
	"git.elenet.me/DA/oncall_list/entity"
)

// QueryTag 查询标签
func (s Service) QueryTag(condition *entity.TagQueryCondition) error {
	return s.tagRepository.TagQuery(condition)
}

// ensureTags
func (s Service) ensureTags(tx *sql.Tx, tagNames []string) ([]int64, error) {
	var err error
	var tagIDs []int64

	dict, err := s.tagRepository.QueryIds(tx, tagNames)
	if err != nil {
		//return nil, fmt.Errorf("ensure: %v", err)
		return nil, err
	}

	for _, tagName := range tagNames {

		id, found := dict[tagName]
		if found {
			tagIDs = append(tagIDs, id)
			continue

		}

		tag := entity.Tag{TagName: tagName, IsDelete: 0}
		err = s.tagRepository.Insert(tx, &tag)
		if err != nil {
			//return nil, fmt.Errorf("tagid: %v", err)
			return nil, err

		}

		tagIDs = append(tagIDs, tag.ID)
	}

	return tagIDs, nil
}
