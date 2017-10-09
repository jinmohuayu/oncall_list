package web

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"git.elenet.me/DA/oncall_list/entity"

	"github.com/martini-contrib/render"
)

// userIndex 用户列表
func (s Server) userIndex(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	r.HTML(http.StatusOK, "index", nil)

}

// userQuery 用户查询
func (s Server) userQuery(request *http.Request, responseWrite http.ResponseWriter, r render.Render) {

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
		Name: sql.NullString{String: strings.Trim(name, " "), Valid: strings.Trim(name, " ") != ""},
		Page: entity.Page{CurrentPageIndex: page, RecordsPerPage: pageSize},
	}

	err = s.service.QueryUser(&condition)
	if err != nil {
		r.JSON(http.StatusInternalServerError, errorResult(err))
		return
	}

	r.JSON(http.StatusOK, successResult(condition))
}

// userDetail 用户详细信息
func (s Server) userDetail(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	id, err := strconv.ParseInt(request.FormValue("id"), 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, errorResult(err))
		return
	}

	user, err := s.service.GetUser(id)
	if err != nil {
		r.JSON(http.StatusInternalServerError, errorResult(err))
		return

	}

	userDetail := entity.UserDetail{
		User:       user,
		Privileges: entity.AllowEdit, // 权限控制，将所有登录用户设置为都可以修改
	}

	r.JSON(http.StatusOK, successResult(userDetail))

}
