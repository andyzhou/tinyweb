package face

import (
	"fmt"
	"github.com/kataras/iris"
	"math"
	"net/url"
	"strings"
)

/*
 * face for web
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//inter macro define
const (
	HttpProtocol = "://"
	MaxPagesPerTime = 10
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

//face info
type Web struct {
}

//construct
func NewWeb() *Web {
	//self init
	this := &Web{
	}
	return this
}

//general page list
func (w *Web) GenPageList(
				req string,
				query string,
				curPage int,
				totalRecords int,
				recPerPage int,
			) (PageInfo, []PageInfo) {
	var (
		startPage int
		endPage int
		active bool
		prevNexPage PageInfo
		pageList = make([]PageInfo, 0)
	)

	if curPage <= 0 {
		curPage = 1
	}

	pages := math.Ceil(float64(totalRecords)/float64(recPerPage))
	totalPagesInt := int(pages)

	//set request and query
	prevNexPage.Request = req
	prevNexPage.Query = query

	//init prev、next page
	if curPage <= 1 {
		prevNexPage.Prev = "1"
		prevNexPage.PrevDisable = true
	}else{
		prevNexPage.Prev = fmt.Sprintf("%d", curPage - 1)
		prevNexPage.PrevDisable = false
	}

	if curPage >= totalPagesInt {
		prevNexPage.Next = fmt.Sprintf("%d", totalPagesInt)
		prevNexPage.NextDisable = true
	}else{
		prevNexPage.Next = fmt.Sprintf("%d", curPage+1)
		prevNexPage.NextDisable = false
	}

	//set batch pages
	startPage = curPage - MaxPagesPerTime/2
	if startPage <= 0 {
		startPage = 1
	}

	endPage = curPage + MaxPagesPerTime/2
	if endPage >= totalPagesInt {
		endPage = totalPagesInt
	}

	//init page list
	for i := startPage; i <= endPage; i++ {
		active = false
		if i == curPage {
			active = true
		}
		pageInfo := PageInfo{
			Idx:fmt.Sprintf("%d", i),
			Query:query,
			Request:req,
			Active:active,
		}
		pageList = append(pageList, pageInfo)
	}
	return prevNexPage, pageList
}


//get refer domain
func (w *Web) GetReferDomain(referUrl string) string {
	var (
		referDomain string
	)
	if referUrl == "" {
		return referDomain
	}
	//find first '://' pos
	protocolLen := len(HttpProtocol)
	protocolPos := strings.Index(referUrl, HttpProtocol)
	if protocolPos <= -1 {
		return referDomain
	}
	//pick domain
	tempBytes := []byte(referUrl)
	tempBytesLen := len(tempBytes)
	prefixLen := protocolPos + protocolLen
	resetUrl := tempBytes[prefixLen:tempBytesLen]
	tempSlice := strings.Split(string(resetUrl), "/")
	if tempSlice == nil || len(tempSlice) <= 0 {
		return referDomain
	}
	referDomain = fmt.Sprintf("%s%s", tempBytes[0:prefixLen], tempSlice[0])
	return referDomain
}

//get general parameter
func (w *Web) GetParameter(
					paraKey string,
					httpForm url.Values,
					ctx iris.Context,
				) string {
	//check and init http form
	if httpForm == nil {
		httpForm = w.GetHttpParameters(ctx)
	}
	//get relate para value
	paraVal := httpForm.Get(paraKey)
	if paraVal == "" {
		paraVal = ctx.Params().Get(paraKey)
		if paraVal == "" {
			paraVal = ctx.PostValueTrim(paraKey)
		}
	}
	return paraVal
}

//get all values of one parameter
func (w *Web) GetParameterValues(
					name string,
					form url.Values,
					ctx iris.Context,
				) []string {
	if form == nil {
		return nil
	}
	vs := form[name]
	if len(vs) == 0 {
		return nil
	}
	return vs
}

//get http request parameters
func (w *Web) GetHttpParameters(ctx iris.Context) url.Values {
	var err error
	//get request uri
	reqUri := w.GetReqUri(ctx)
	if reqUri == "" {
		return nil
	}
	//parse form
	queryForm, err := url.ParseQuery(reqUri)
	if err != nil {
		return nil
	}
	return queryForm
}

//get request uri
func (w *Web) GetReqUri(ctx iris.Context) string {
	var (
		reqUriFinal string
	)

	reqUri := ctx.Request().URL.RawQuery
	reqUriNew, err := url.QueryUnescape(reqUri)
	if err != nil {
		return reqUriFinal
	}
	reqUriFinal = reqUriNew
	return reqUriFinal
}

//get client ip
func (w *Web) GetClientIp(ctx iris.Context) string {
	clientIp := ctx.RemoteAddr()
	xRealIp := ctx.GetHeader("X-Real-IP")
	xForwardedFor := ctx.GetHeader("X-Forwarded-For")
	if clientIp != "" {
		return clientIp
	}else{
		if xRealIp != "" {
			clientIp = xRealIp
		}else{
			clientIp = xForwardedFor
		}
	}
	return clientIp
}
