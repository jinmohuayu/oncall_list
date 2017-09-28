package web

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"git.elenet.me/DA/oncall_list/logic"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	staticPath     = "/static"
	stateChecked   = "on"
	idSplitter     = "|"
	recordsPerPage = 20
	errorPagePath  = "/error"
)

var (
	// ErrUserSignOut 用户未登录
	ErrUserSignOut = errors.New("用户未登录")
	// ErrInvalidSignIn 非法的登录请求
	ErrInvalidSignIn = errors.New("非法的登录请求")
	// ErrBadRequest 请求参数不正确
	ErrBadRequest = errors.New("请求参数不正确")
)

// ServerConfig Web服务器配置
type ServerConfig struct {
	Port int     // 监听端口
	DB   *sql.DB // 数据库连接
}

// Server Web服务器
type Server struct {
	config  *ServerConfig // 配置
	logger  *log.Logger   // 日志
	service *logic.Service
}

// NewServer 新建Web服务器
func NewServer(config *ServerConfig, logger *log.Logger) *Server {
	return &Server{
		config:  config,
		logger:  logger,
		service: logic.NewService(logic.ServiceConfig{DB: config.DB}, logger),
	}
}

// StartAndWait 启动监听服务
func (s Server) StartAndWait() {

	m := martini.New()

	// 异常显示
	m.Use(martini.Recovery())

	// 静态路径 /static
	m.Use(martini.Static("web/static", martini.StaticOptions{
		Prefix:      staticPath,
		SkipLogging: true,
	}))

	// 模板配置
	m.Use(render.Renderer(render.Options{
		Directory: "server/static/html", // Specify what path to load the templates from.
		// Layout:          "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions:      []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:         "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON:      false,                      // 输出紧凑JSON
		IndentXML:       true,                       // 输出可读的XML
		HTMLContentType: "text/html",                // HTML输出格式
		Funcs: []template.FuncMap{{
			"rowindex": displayRowIndex,
		}},
		// Specify helper function maps for templates to access.
		// Delims: render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
	}))

	// 路由
	r := martini.NewRouter()
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	// store := sessions.NewCookieStore([]byte(cookieStoreKey))
	// store.Options(sessions.Options{
	// 	Path:   "/",
	// 	Domain: cookieDomain,
	// 	MaxAge: 43200, // 12hours
	// })
	// m.Use(sessions.Sessions(sessionName, store))

	// 登录验证
	m.Use(s.signInValidation)

	cm := &martini.ClassicMartini{Martini: m, Router: r}

	// 注册路由
	s.registerRoute(cm)

	// 监听
	address := fmt.Sprintf(":%d", s.config.Port)
	s.logger.Printf("oncall web服务已启动，端口%s", address)
	cm.RunOnAddr(address)
}

// error 错误页
func (s Server) error(request *http.Request, r render.Render) {
	message, _ := url.QueryUnescape(request.FormValue("m"))
	r.HTML(http.StatusOK, "error", message)
}

// redirectToErrorPage 跳转到错误提示页
func (s Server) redirectToErrorPage(request *http.Request, responseWriter http.ResponseWriter, err error) {

	message := fmt.Sprint(err)
	if err == sql.ErrNoRows {
		message = "记录不存在"
	}

	redirectURL := fmt.Sprintf("%s?m=%s", errorPagePath, url.QueryEscape(message))
	http.Redirect(responseWriter, request, redirectURL, http.StatusTemporaryRedirect)
}

// displayRowIndex 行显示序号主要用于html页面分页结果表格的记录序号显示
func displayRowIndex(currentPageIndex, pageSize, index int) string {
	return strconv.Itoa((currentPageIndex-1)*pageSize + index + 1)
}
