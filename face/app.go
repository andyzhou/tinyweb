package face

import (
	"fmt"
	"github.com/andyzhou/tinyweb/define"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	"net/http"
)

/*
 * gin app face
 * - github.com/gin-gonic/gin
 */

//inter macro define
const (
	ViewExtName = ".html"
)

//face info
type WebApp struct {
	port int //web port
	httpServer *http.Server
	server *gin.Engine //gin server
	tplPath string //tpl root path
	tpl iface.ITpl //tpl interface
	runner *iris.Runner //iris runner
}

//construct
func NewWebApp(g *gin.Engine) *WebApp {
	var (
		s *gin.Engine
	)
	//check
	if g != nil {
		s = g
	}else{
		s = gin.Default()
	}
	//self init
	this := &WebApp{
		server: s,
	}
	return this
}

//start app service
func (f *WebApp) Start(port int) bool {
	if port <= 0 {
		return false
	}

	//set port and address
	f.port = port
	addr := fmt.Sprintf(":%v", port)

	//start app
	f.server.Run(addr)
	return true
}

//register root app entry
//url like: /xxx or /xxx/{ParaName:string|integer}
func (f *WebApp) RegisterRootApp(
						rootUrlPara string,
						face iface.IWebSubApp,
					) bool {
	//check
	if rootUrlPara == "" || face == nil {
		return false
	}

	//root request route
	rootAnyPath := fmt.Sprintf("/%v/*%v", rootUrlPara, define.AnyPath)

	//set get、post request
	f.server.Any(rootAnyPath, face.Entry)
	return true
}

//get tpl interface
func (f *WebApp) GetTplInterface() iface.ITpl {
	return f.tpl
}

//set error code and cb
func (f *WebApp) SetErrCode(code int, cb func(c *gin.Context)) bool {
	if code < 0 || cb == nil {
		return false
	}
	return true
}

//set static file path
func (f *WebApp) SetStaticPath(url, path string) bool {
	if url == "" || path == "" {
		return false
	}
	f.server.Static(url, path)
	return true
}

//set tpl root path
func (f *WebApp) SetTplPath(path string) bool {
	if path == "" {
		return false
	}
	//set tpl path
	f.tplPath = path

	//init templates
	//tpl := iris.HTML(f.tplPath, ViewExtName).Reload(true)
	//f.app.RegisterView(tpl)
	f.server.LoadHTMLGlob(fmt.Sprintf("%v/*", f.tplPath))

	//init tpl face
	f.tpl = NewTpl()
	return true
}