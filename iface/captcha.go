package iface

import (
	"github.com/gin-gonic/gin"
)

type ICaptcha interface {
	GetCookie(c *gin.Context) (string, error)
	GenImg(c *gin.Context)
	SetCookiePara(name string, expire int) bool
	SetCaptchaNum(num int) bool
	SetSize(width, height int) bool
	SetFontPath(path string) bool
}