package face

import (
	"fmt"
	"github.com/andyzhou/tinyweb/define"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/gin-gonic/gin"
	"sync"
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
	server *gin.Engine //gin server
	tplPath string //tpl root path
	tpl iface.ITpl //tpl interface
	//runner *iris.Runner //iris runner
	wg sync.WaitGroup
}

//construct
func NewWebApp(g ...*gin.Engine) *WebApp {
	var (
		s *gin.Engine
	)
	//check
	if g != nil && len(g) > 0 {
		s = g[0]
	}else{
		s = gin.Default()
	}
	//self init
	this := &WebApp{
		server: s,
	}
	return this
}

//stop app
func (f *WebApp) Stop() {
	f.wg.Done()
}

//start app
func (f *WebApp) Start(port int) bool {
	if port <= 0 {
		return false
	}

	//set port and address
	f.port = port
	addr := fmt.Sprintf(":%v", port)

	//start app
	f.wg.Add(1)
	f.server.Run(addr)
	f.wg.Wait()
	return true
}

//register root app entry
//url like: /xxx or /xxx/{ParaName:string|integer}
func (f *WebApp) RegisterSubApp(
						reqUrlPara string,
						face iface.IWebSubApp,
						incPathParas ...bool,
					) bool {
	var (
		requestAnyPath string
		incPathPara bool
	)
	//check
	if reqUrlPara == "" || face == nil {
		return false
	}
	if incPathParas != nil && len(incPathParas) > 0 {
		incPathPara = incPathParas[0]
	}

	//root request route
	if incPathPara {
		requestAnyPath = reqUrlPara
	}else{
		requestAnyPath = fmt.Sprintf("/%v/*%v", reqUrlPara, define.AnyPath)
	}

	//set get、post request
	f.server.Any(requestAnyPath, face.Entry)
	return true
}

//get tpl interface
func (f *WebApp) GetTplInterface() iface.ITpl {
	return f.tpl
}

//get gin engine
func (f *WebApp) GetGin() *gin.Engine {
	return f.server
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
	f.server.LoadHTMLGlob(fmt.Sprintf("%v/*", f.tplPath))

	//init tpl face
	f.tpl = NewTpl()
	return true
}

//set gin reference
func (f *WebApp) SetGin(gin *gin.Engine) {
	if gin == nil {
		return
	}
	f.server = gin
}