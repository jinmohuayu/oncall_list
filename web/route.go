package web

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

//  registerRoute 注册路由
func (s Server) registerRoute(m *martini.ClassicMartini) {

	// 用户
	m.Get("/users", s.userIndex)
	m.Get("/api/users", s.userQuery)
	m.Get("/api/users/details", s.userDetail)
	m.Post("/api/users/updateTag", binding.Json(UpdateUserTagForm{}), s.userUpdateTag)

	// 标签
	m.Get("/api/tags", s.tagIndex)

	// 错误页
	m.Get("/error", s.error)

}
