package face

import (
	"fmt"
	"github.com/andyzhou/tinyweb/define"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/kataras/iris/v12"
)

/*
 * iris app face
 * - github.com/kataras/iris/v12
 */

//inter macro define
const (
	ViewExtName = ".html"
)

//face info
type IrisApp struct {
	tplPath string //tpl root path
	tpl iface.ITpl //tpl interface
	app *iris.Application //iris application instance
	runner *iris.Runner //iris runner
}

//construct
func NewIrisApp() *IrisApp {
	this := &IrisApp{
		app: iris.New(),
	}
	return this
}

//start app service
func (f *IrisApp) Start(port int) bool {
	if port <= 0 {
		return false
	}

	//init runner
	runner := iris.Addr(fmt.Sprintf(":%d", port))
	f.runner = &runner

	//start app
	go f.app.Run(*f.runner)
	return true
}

//register root app entry
func (f *IrisApp) RegisterRootApp(
						rootUrlPara string,
						face iface.IIrisSubApp,
						methods ...string,
					) bool {
	var (
		reqMethod string
	)

	//check
	if rootUrlPara == "" || face == nil {
		return false
	}

	//get request method
	if methods == nil {
		reqMethod = define.ReqMethodOfGet
	}else{
		reqMethod = methods[0]
	}

	//set relate request
	switch reqMethod {
	case define.ReqMethodOfPost://post
		f.app.Post(rootUrlPara, face.Entry)
	case define.ReqMethodOfGet://get
		fallthrough
	default:
		f.app.Get(rootUrlPara, face.Entry)
	}
	return true
}

//get tpl interface
func (f *IrisApp) GetTplInterface() iface.ITpl {
	return f.tpl
}

//set error code and cb
func (f *IrisApp) SetErrCode(code int, cb func(ctx iris.Context)) bool {
	if code < 0 || cb == nil {
		return false
	}
	f.app.OnErrorCode(code, cb)
	return true
}

//set static file path
func (f *IrisApp) SetStaticPath(url, path string) bool {
	if url == "" || path == "" {
		return false
	}
	f.app.HandleDir(url, path)
	return true
}

//set tpl root path
func (f *IrisApp) SetTplPath(path string) bool {
	if path == "" {
		return false
	}
	//set tpl path
	f.tplPath = path

	//init templates
	tpl := iris.HTML(f.tplPath, ViewExtName)
	f.app.RegisterView(tpl)

	//init tpl face
	f.tpl = NewTpl(tpl)
	return true
}