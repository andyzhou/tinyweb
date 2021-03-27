package iface

import (
	"github.com/kataras/iris/v12"
	"net/url"
)

//single page info
type PageInfo struct {
	Idx string //page num
	Active bool //current page
	Prev string
	PrevDisable bool //disable click
	Next string
	NextDisable bool //disable click
	Query string `main query param info`
	Request string `main request`
}

type IWeb interface {
	GenPageList(
		req string,
		query string,
		curPage int,
		totalRecords int,
		recPerPage int,
	) (PageInfo, []PageInfo)
	TrimHtml(src string, needLower bool) string
	SubString(source string, start int, length int) string
	GetReferDomain(referUrl string) string
	GetParameter(paraKey string, httpForm url.Values, ctx iris.Context) string
	GetParameterValues(name string, form url.Values, ctx iris.Context) []string
	GetHttpParameters(ctx iris.Context) url.Values
	GetReqUri(ctx iris.Context) string
	GetClientIp(ctx iris.Context) string
}