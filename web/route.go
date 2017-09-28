package web

import "github.com/go-martini/martini"

//  registerRoute 注册路由
func (s Server) registerRoute(m *martini.ClassicMartini) {

	// 首页
	m.Get("/users", s.userIndex)
	// 错误页
	m.Get("/error", s.error)

}
