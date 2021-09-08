package iface

import "github.com/kataras/iris/v12"

type ICaptcha interface {
	GenImg(ctx iris.Context)
	SetCookiePara(name string, expire int) bool
	SetCaptchaNum(num int) bool
	SetSize(width, height int) bool
	SetFontPath(path string) bool
}