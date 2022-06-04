package face

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"sync"

	//"github.com/kataras/iris/v12"
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

//global variable for single instance
var (
	_web *Web
	_webOnce sync.Once
)

//face info
type Web struct {
}

//get single instance
func GetWeb() *Web {
	_webOnce.Do(func() {
		_web = NewWeb()
	})
	return _web
}

//construct
func NewWeb() *Web {
	//self init
	this := &Web{
	}
	return this
}

//download data as file
func (f *Web) DownloadAsFile(downloadName string, data []byte, c *gin.Context) error {
	//check
	if downloadName == "" || data == nil {
		return errors.New("invalid parameter")
	}

	//setup header
	c.Writer.Header().Add("Content-type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename= " + downloadName)

	//write data into download file
	_, err := c.Writer.Write(data)
	return err
}

//calculate total pages
func (f *Web) CalTotalPages(total, size int) int {
	return int(math.Ceil(float64(total) / float64(size)))
}

//get json request body
func (f *Web) GetJsonRequest(c *gin.Context, obj interface{}) error {
	//try read body
	jsonByte, err := f.GetRequestBody(c)
	if err != nil {
		return err
	}
	//try decode json data
	err = json.Unmarshal(jsonByte, obj)
	return err
}

//get request body
func (f *Web) GetRequestBody(c *gin.Context) ([]byte, error) {
	return ioutil.ReadAll(c.Request.Body)
}

//get request para
func (f *Web) GetPara(name string, c *gin.Context) string {
	//decode request url
	decodedReqUrl, _ := url.PathUnescape(c.Request.URL.RawQuery)
	values, _ := url.ParseQuery(decodedReqUrl)

	//get act from url
	if values != nil {
		paraVal := values.Get(name)
		if paraVal != "" {
			return paraVal
		}
	}

	//get act from query, post.
	paraVal := c.Query(name)
	if paraVal == "" {
		//get from post
		paraVal = c.PostForm(name)
	}
	return paraVal
}

//general page list
func (f *Web) GenPageList(
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
func (f *Web) SubString(source string, start int, length int) string {
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
func (f *Web) TrimHtml(src string, needLower bool) string {
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
func (f *Web) GetReferDomain(referUrl string) string {
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

//get request uri
func (f *Web) GetReqUri(ctx *gin.Context) string {
	var (
		reqUriFinal string
	)
	reqUri := ctx.Request.URL.RawQuery
	reqUriNew, err := url.QueryUnescape(reqUri)
	if err != nil {
		return reqUriFinal
	}
	reqUriFinal = reqUriNew
	return reqUriFinal
}

//get client ip
func (f *Web) GetClientIp(ctx *gin.Context) string {
	clientIp := ctx.Request.RemoteAddr
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
