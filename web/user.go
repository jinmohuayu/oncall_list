package web

import (
	"database/sql"
	"errors"
	"git.elenet.me/DA/oncall_list/entity"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"strings"
)

// userIndex 用户列表
func (s Server) userIndex(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	r.HTML(http.StatusOK, "index", nil)

}

// userQuery 用户查询
func (s Server) userQuery(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	name := request.FormValue("name")
	tagIDs := request.FormValue("tagids")
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

	if strings.Trim(tagIDs, " ") != "" {
		condition.TagIds = strings.Split(tagIDs, ",")
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

// UserUpdateTag 更新当前用户的Tag
func (s Server) userUpdateTag(form UpdateUserTagForm, request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	s.logger.Printf("%#v", form)
	if form.UserID == 0 {
		r.JSON(http.StatusBadRequest, errorResult(errors.New("userid can't be null")))
		return

	}

	tagNames := strings.Split(form.TagNames, ",")
	var names []string
	for _, name := range tagNames {
		if name == "" {
			continue
		}
		names = append(names, name)
	}

	err := s.service.UpdateUserTag(form.UserID, names)
	if err != nil {
		r.JSON(http.StatusInternalServerError, errorResult(err))
		return

	}

	r.JSON(http.StatusOK, successResult(nil))
}
