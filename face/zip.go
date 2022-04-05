package face

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"sync"
)

/*
 * face for data zip
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//global variable for single instance
var (
	_zip *Zip
	_zipOnce sync.Once
)

//face info
type Zip struct {
}

//get single instance
func GetZip() *Zip {
	_zipOnce.Do(func() {
		_zip = NewZip()
	})
	return _zip
}

//construct
func NewZip() *Zip {
	//self init
	this := &Zip{
	}
	return this
}

//un compress data
func (f *Zip) UnZipData(resp []byte) (bool, map[string]interface{}) {
	var (
		outBytes bytes.Buffer
	)

	//base64 decode
	respStr := string(resp)
	byteData, err := base64.StdEncoding.DecodeString(respStr)
	if err != nil {
		return false, nil
	}

	//zip un compress
	byteReader := bytes.NewReader(byteData)
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		return false, nil
	}

	_, err = io.Copy(&outBytes, gzipReader)
	if err != nil {
		return false, nil
	}

	//json decode
	result := make(map[string]interface{})
	err = json.Unmarshal(outBytes.Bytes(), &result)
	if err != nil {
		return false, nil
	}

	return true, result
}

//compress data
func (f *Zip) ZipData(req map[string]interface{}) (bool, []byte) {
	var (
		in bytes.Buffer
	)
	//basic check
	if req == nil || len(req) <= 0 {
		return false, nil
	}

	//json encode
	jsonByte, err := json.Marshal(req)
	if err != nil {
		return false, nil
	}

	//zip compress
	zipWriter := gzip.NewWriter(&in)
	_, err = zipWriter.Write(jsonByte)
	if err != nil {
		zipWriter.Close()
		return false, nil
	}
	err = zipWriter.Close()
	if err != nil {
		log.Println("err:", err.Error())
	}

	//base64 encode
	out := make([]byte, 2048)
	base64.StdEncoding.Encode(out, in.Bytes())

	return true, out
}