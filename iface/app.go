package iface

import (
	"github.com/gin-gonic/gin"
)

//web app interface
type IWebApp interface {
	Stop()
	Start(port int) bool
	RegisterSubApp(reqUrlPara string, face IWebSubApp, incPathParas ...bool) bool

	//get
	GetTplInterface() ITpl
	GetGin() *gin.Engine

	//set
	SetStaticPath(url, path string) bool
	SetTplPath(path string) bool
	SetGin(gin *gin.Engine)
}

//sub web app interface
type IWebSubApp interface {
	Entry(c *gin.Context)
}