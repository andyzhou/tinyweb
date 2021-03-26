package iface

import (
	"github.com/kataras/iris"
	"net/url"
)

type IWeb interface {
	GetReferDomain(referUrl string) string
	GetParameter(paraKey string, httpForm url.Values, ctx iris.Context) string
	GetParameterValues(name string, form url.Values, ctx iris.Context) []string
	GetHttpParameters(ctx iris.Context) url.Values
	GetReqUri(ctx iris.Context) string
	GetClientIp(ctx iris.Context) string
}