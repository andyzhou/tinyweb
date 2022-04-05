package face

import (
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
 * face for tpl, base on iris
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//inter macro define
const (
	TimeLayOut = "2006-01-02 15:04:05" //can't be changed!!!
	SubStrMaxLen = 10
)

//face info
type Tpl struct {
	funcMap template.FuncMap
	//tpl *view.HTMLEngine
}

//construct
func NewTpl() *Tpl {
	self := &Tpl{
		funcMap: template.FuncMap{},
	}
	return self
}

//get func map
func (f *Tpl) GetFuncMap() template.FuncMap {
	return f.funcMap
}

//register diy tpl func
func (f *Tpl) RegisterTplFunc(
					name string,
					cb interface{},
				) bool {
	if name == "" || cb == nil {
		return false
	}
	if f.funcMap == nil {
		return false
	}
	//check
	_, ok := f.funcMap[name]
	if ok {
		return false
	}
	f.funcMap[name] = cb
	return true
}

//register base tpl func
func (f *Tpl) RegisterTplBaseFunc() bool {
	if f.funcMap == nil {
		return false
	}

	//add html function
	f.funcMap["html"] = f.HtmlData

	//add substring function
	f.funcMap["substr"] = f.SubStr

	//trim html function
	f.funcMap["trimHtml"] = f.TrimHtml

	//remove high light mark function
	f.funcMap["removeMark"] = f.RemoveMark

	//time stamp format
	f.funcMap["date"] = f.TimeStamp2Date
	f.funcMap["datetime"] = f.TimeStamp2DateTime
	f.funcMap["dayTime"] = f.TimeStampToDayStr
	f.funcMap["second2Time"] = f.Seconds2TimeStr

	return true
}

///////////////
//private func
///////////////

//remove mark
func (f *Tpl) RemoveMark(text string) string {
	if text == "" {
		return text
	}
	re, _ := regexp.Compile("\\<\\/?mark\\>")
	text = re.ReplaceAllString(text, "")
	return text
}

//html data
func (f *Tpl) HtmlData(text string) template.HTML {
	return template.HTML(text)
}

//sub string
func (f *Tpl) SubStr(text string) string {
	if len(text) <= SubStrMaxLen {
		return text
	}
	final := fmt.Sprintf("%s..", f.SubString(text, 0, SubStrMaxLen))
	return final
}

//convert timestamp to date format, like YYYY-MM-DD
func (f *Tpl) TimeStamp2Date(timeStamp int64) string {
	dateTime := time.Unix(timeStamp, 0).Format(TimeLayOut)
	tempSlice := strings.Split(dateTime, " ")
	if tempSlice == nil || len(tempSlice) <= 0 {
		return ""
	}
	return tempSlice[0]
}

//convert timestamp to data time string format
func (f *Tpl) TimeStamp2DateTime(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format(TimeLayOut)
}

//convert timestamp like 'Oct 10, 2020' format
func (f *Tpl) TimeStampToDayStr(timeStamp int64) string {
	date := f.TimeStamp2Date(timeStamp)
	if date == "" {
		return  ""
	}
	tempSlice := strings.Split(date, "-")
	if tempSlice == nil || len(tempSlice) < 3 {
		return ""
	}
	year := tempSlice[0]
	month, _ := strconv.Atoi(tempSlice[1])
	day := tempSlice[2]
	return fmt.Sprintf("%s %s, %s", time.Month(month).String(), day, year)
}

//convert seconds to time string format
func (f *Tpl) Seconds2TimeStr(seconds int) string {
	var (
		hourStr, minuteStr, secondStr string
	)

	if seconds <= 0 {
		return ""
	}

	hourInt := seconds / 3600
	minuteInt := (seconds - hourInt * 3600) / 60
	secondInt := seconds - hourInt * 3600 - minuteInt * 60

	if hourInt > 0 {
		if hourInt > 9 {
			hourStr = fmt.Sprintf("%d:", hourInt)
		}else{
			hourStr = fmt.Sprintf("0%d:", hourInt)
		}
	}

	if minuteInt > 9 {
		minuteStr = fmt.Sprintf("%d", minuteInt)
	}else{
		minuteStr = fmt.Sprintf("0%d", minuteInt)
	}

	if secondInt > 9 {
		secondStr = fmt.Sprintf("%d", secondInt)
	}else{
		secondStr = fmt.Sprintf("0%d", secondInt)
	}

	//format time string
	timeStr := fmt.Sprintf("%s%s:%s", hourStr, minuteStr, secondStr)
	return timeStr
}

//sub string, support utf8 string
func (f *Tpl) SubString(source string, start int, length int) string {
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
func (f *Tpl) TrimHtml(src string) string {
	var (
		re *regexp.Regexp
	)

	if src == "" {
		return src
	}

	//convert to lower
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//remove style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//remove script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	return strings.TrimSpace(src)
}

