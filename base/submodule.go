package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
	"strconv"
)

/*
 * base module
 * - shared inter api
 */

//face info
type SubModule struct {
}

//download data as file
func (f *SubModule) DownloadAsFile(downloadName string, data []byte, c *gin.Context) error {
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
func (f *SubModule) CalTotalPages(total, size int) int {
	return int(math.Ceil(float64(total) / float64(size)))
}

//get para value
func (f *SubModule) GetPara(name string, c *gin.Context) string {
	//get act from query, post.
	act := c.Query(name)
	if act == "" {
		//get from post
		act = c.PostForm(name)
	}
	return act
}

//get request body
func (f *SubModule) GetRequestBody(c *gin.Context) ([]byte, error) {
	return ioutil.ReadAll(c.Request.Body)
}

//get path param
func (f *SubModule) GetPathPara(para string, c *gin.Context) string {
	return c.Param(para)
}

//convert string to int64
func (f *SubModule) Str2Int(val string) int64 {
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return intVal
}