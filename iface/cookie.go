package iface

import (
	"context"
	"github.com/kataras/iris"
)

type ICookie interface {
	DelCookie(key string, ctx context.Context) bool
	GetCookie(key string, ctx context.Context) string
	SetCookie(key, value string, totalSeconds int, ctx context.Context) bool
	SetCoreData(app *iris.Application) bool
}
