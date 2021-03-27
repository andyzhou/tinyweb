package iface

import "github.com/kataras/iris/v12"

type ICaptcha interface {
	GenImg(ctx iris.Context)
	SetCookiePara(name string, expire int) bool
}