package iface

import (
	"github.com/gin-gonic/gin"
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
	DownloadAsFile(downloadName string, data []byte, c *gin.Context) error
	CalTotalPages(total, size int) int
	GetJsonRequest(c *gin.Context, obj interface{}) error
	GetRequestBody(c *gin.Context) ([]byte, error)
	GetPara(name string, c *gin.Context) string
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
	GetReqUri(ctx *gin.Context) string
	GetClientIp(ctx *gin.Context) string
}