package iface

import (
	"github.com/gin-gonic/gin"
)

//web app interface
type IWebApp interface {
	Start(port int) bool
	RegisterRootApp(rootUrlPara string, face IWebSubApp) bool
	GetTplInterface() ITpl
	SetErrCode(code int, cb func(c *gin.Context)) bool
	SetStaticPath(url, path string) bool
	SetTplPath(path string) bool
}

//sub web app interface
type IWebSubApp interface {
	Entry(c *gin.Context)
}