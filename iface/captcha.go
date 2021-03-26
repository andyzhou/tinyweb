package iface

import "github.com/kataras/iris"

type ICaptcha interface {
	GenImg(ctx iris.Context)
	SetCookie(name string, expire int64) bool
}