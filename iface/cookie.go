package iface

import (
	"github.com/kataras/iris/v12"
)

type ICookie interface {
	DelCookie(key string, ctx iris.Context) bool
	GetCookie(key string, ctx iris.Context) string
	SetCookie(key, value string, totalSeconds int, ctx iris.Context) bool
	SetCoreData(app *iris.Application) bool
}
