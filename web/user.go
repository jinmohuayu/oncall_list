package web

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"git.elenet.me/DA/oncall_list/entity"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

// userIndex 用户列表
func (s Server) userIndex(request *http.Request, responseWriter http.ResponseWriter, session sessions.Session, r render.Render) {

	name := request.FormValue("name")
	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(request.FormValue("pageSize"))
	if err != nil {
		pageSize = recordsPerPage
	}

	condition := entity.UserQueryCondition{
		Name: sql.NullString{String: name, Valid: strings.Trim(name, " ") != ""},
		Page: entity.Page{CurrentPageIndex: page, RecordsPerPage: pageSize},
	}

	err = s.service.QueryUser(&condition)
	if err != nil {
		r.JSON(http.StatusOK, errorResult(err))
		return
	}

	users := condition.Records.([]entity.User)

	r.JSON(http.StatusOK, successResult(users))
}
