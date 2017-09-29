package web

import "github.com/go-martini/martini"

//  registerRoute 注册路由
func (s Server) registerRoute(m *martini.ClassicMartini) {

	// 用户
	m.Get("/users", s.userIndex)
	m.Get("/api/users", s.userQuery)

	// 标签
	m.Get("/api/tags", s.tagIndex)

	// 错误页
	m.Get("/error", s.error)

}
