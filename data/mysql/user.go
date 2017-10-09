package mysql

import (
	"bytes"
	"database/sql"
	"fmt"

	"git.elenet.me/DA/oncall_list/entity"
)

// UserRepository 用户Repository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 新建用户Repository
func NewUserRepository(db *sql.DB) *UserRepository {

	return &UserRepository{db: db}
}

// GetByID 获取指定ID的模块
func (r UserRepository) GetByID(id int64) (*entity.User, error) {

	module := new(entity.User)
	row := r.db.QueryRow("select U.id, U.name, U.department, U.product, U.email, U.phone_num, U.remark, U.is_delete, U.created_at, U.updated_at from user_info U where id = ? ", id)
	err := module.Scan(row)

	return module, err
}

// Query 查询
func (r UserRepository) Query(condition *entity.UserQueryCondition) error {

	var conditionSQL bytes.Buffer
	var parameters []interface{}

	if condition.Name.Valid {
		conditionSQL.WriteString("and instr(U.name, ?) > 0 ")
		parameters = append(parameters, condition.Name.String)
	}

	listSQL := bytes.NewBufferString("select U.id, U.name, U.department, U.product, U.email, U.phone_num, U.remark, U.is_delete, U.created_at, U.updated_at from user_info U where 1=1 ")
	listSQL.WriteString(conditionSQL.String())
	listSQL.WriteString("order by U.id asc ")
	if condition.Page.CurrentPageIndex > 0 {
		listSQL.WriteString(fmt.Sprintf("limit %d, %d", (condition.Page.CurrentPageIndex-1)*condition.Page.RecordsPerPage, condition.Page.RecordsPerPage))
	}

	fmt.Printf(listSQL.String())
	rows, err := r.db.Query(listSQL.String(), parameters...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		user := entity.User{}
		err = user.ScanRows(rows)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	countSQL := bytes.NewBufferString("select count(0) from user_info U where 1=1 ")
	countSQL.WriteString(conditionSQL.String())

	var count int
	err = r.db.QueryRow(countSQL.String(), parameters...).Scan(&count)
	if err != nil {
		return err
	}

	condition.Page.TotalRecords = count
	condition.Page.Records = users

	return nil
}
