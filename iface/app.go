package iface

import "github.com/kataras/iris/v12"

//iris app interface
type IIrisApp interface {
	Start(port int) bool
	RegisterRootApp(rootUrlPara string, face IIrisSubApp) bool
	GetTplInterface() ITpl
	SetErrCode(code int, cb func(ctx iris.Context)) bool
	SetStaticPath(url, path string) bool
	SetTplPath(path string) bool
}

//sub iris app interface
type IIrisSubApp interface {
	PostEntry(ctx iris.Context)
	GetEntry(ctx iris.Context)
}