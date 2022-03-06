package face

import (
	"fmt"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/kataras/iris/v12"
	"math"
	"net/url"
	"regexp"
	"strings"
)

/*
 * face for web, base on iris
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//inter macro define
const (
	HttpProtocol = "://"
	MaxPagesPerTime = 10
)



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
			) (iface.PageInfo, []iface.PageInfo) {
	var (
		startPage int
		endPage int
		active bool
		prevNexPage iface.PageInfo
		pageList = make([]iface.PageInfo, 0)
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
		pageInfo := iface.PageInfo{
			Idx:fmt.Sprintf("%d", i),
			Query:query,
			Request:req,
			Active:active,
		}
		pageList = append(pageList, pageInfo)
	}
	return prevNexPage, pageList
}

//sub string, support utf8 string
func (w *Web) SubString(source string, start int, length int) string {
	rs := []rune(source)
	len := len(rs)
	if start < 0 {
		start = 0
	}
	if start >= len {
		start = len
	}
	end := start + length
	if end > len {
		end = len
	}
	return string(rs[start:end])
}

//remove html tags
func (w *Web) TrimHtml(src string, needLower bool) string {
	var (
		re *regexp.Regexp
	)

	if needLower {
		//convert to lower
		re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		src = re.ReplaceAllStringFunc(src, strings.ToLower)
	}

	//remove style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//remove script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	return strings.TrimSpace(src)
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
