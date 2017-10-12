package web

import (
	"github.com/martini-contrib/render"
	"net/http"
)

// signInValidation 登录验证
func (s Server) signInValidation(request *http.Request, responseWriter http.ResponseWriter, r render.Render) {

	// ignorePrefix := []string{staticPath, signInPath, signOutPath, errorPagePath, ldapSignInPath, suningSignInPath}
	// for _, prefix := range ignorePrefix {
	// 	if strings.HasPrefix(request.RequestURI, prefix) {
	// 		return
	// 	}
	// }

	// // 判断登录
	// _, err := s.currentUser(session)
	// if err != nil {
	// 	if err != ErrUserSignOut {
	// 		s.logger.Errorf("登录验证时发生错误: %v", err)
	// 		s.redirectToErrorPage(request, responseWriter, err)
	// 		return
	// 	}

	// 	// 未登录的跳转到登录页
	// 	redirectURL := fmt.Sprintf("%s?referer=http://%s%s", signInPath, request.Host, request.RequestURI)
	// 	// s.logger.Warmf("用户未登录，跳转到登录页:%v", redirectURL)

	// 	http.Redirect(responseWriter, request, redirectURL, http.StatusTemporaryRedirect)

	// 	return
	// }
}
