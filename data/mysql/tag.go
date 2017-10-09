package mysql

import (
	"bytes"
	"database/sql"
	"fmt"
	"git.elenet.me/DA/oncall_list/entity"
)

// TagRepository 标签Repository
type TagRepository struct {
	db *sql.DB
}

// NewTagRepository 新建标签Repository
func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db: db}
}

// tagQuery 根据传过来的查询条件进行查询
func (r TagRepository) TagQuery(condition *entity.TagQueryCondition) error {

	var conditionSQL bytes.Buffer
	var parameters []interface{}

	// TagID 传入的是tagID
	if condition.TagID.Valid {
		conditionSQL.WriteString("and instr(T.id, ?) > 0 ")
		parameters = append(parameters, condition.TagID.Valid)
	}

	// TagName 传入的是TagName
	if condition.TagName.Valid {
		conditionSQL.WriteString("and instr(T.tag_name, ?) > 0 ")
		parameters = append(parameters, condition.TagName.String)
	}

	// UserID 传入的是指定用户
	if condition.UserID.Valid {
		conditionSQL.WriteString("and exists(select 1 from user_tag UT where UT.tag_id = T.id and UT.user_id = ?) ")
		parameters = append(parameters, condition.UserID.Int64)
	}

	listSQL := bytes.NewBufferString(`select T.id, T.tag_name, T.is_delete, T.created_at, T.updated_at from tag T where 1=1 `)
	listSQL.WriteString(conditionSQL.String())
	listSQL.WriteString("order by T.id asc ")
	if condition.Page.CurrentPageIndex > 0 {
		listSQL.WriteString(fmt.Sprintf("limit %d, %d", (condition.Page.CurrentPageIndex-1)*condition.Page.RecordsPerPage, condition.Page.RecordsPerPage))
	}

	// 查询SQL,rows为查询的结果
	rows, err := r.db.Query(listSQL.String(), parameters...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Rows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows:
	var tags []entity.Tag
	for rows.Next() {
		tag := entity.Tag{}
		err = tag.ScanRows(rows)
		if err != nil {
			return err
		}

		tags = append(tags, tag)
	}

	countSQL := bytes.NewBufferString("select count(0) from tag T where 1=1 ")
	countSQL.WriteString(conditionSQL.String())

	var count int
	err = r.db.QueryRow(countSQL.String(), parameters...).Scan(&count)
	if err != nil {
		return err
	}

	condition.Page.TotalRecords = count
	condition.Page.Records = tags

	return nil
}
