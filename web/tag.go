package web

import (
	"database/sql"
	"git.elenet.me/DA/oncall_list/entity"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"strings"
)

// tagIndex 用户索引tags列表
func (s Server) tagIndex(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	tagID := request.FormValue("tagid")
	tagName := request.FormValue("tagname")

	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(request.FormValue("pageSize"))
	if err != nil {
		pageSize = recordsPerPage
	}

	conditionTag := entity.TagQueryCondition{
		TagID:   sql.NullString{String: tagID, Valid: strings.Trim(tagID, " ") != ""},
		TagName: sql.NullString{String: tagName, Valid: strings.Trim(tagName, " ") != ""},
		Page:    entity.Page{CurrentPageIndex: page, RecordsPerPage: pageSize},
	}

	err = s.service.QueryTag(&conditionTag)
	if err != nil {
		return
	}

	tags := conditionTag.Records.([]entity.Tag)
	r.JSON(http.StatusOK, successResult(tags))
}
