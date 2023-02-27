package face

import (
	"errors"
	"github.com/andyzhou/tinyweb/iface"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"sync"
)

/*
 * face for cookie
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//internal macro defines
const (
	CookieHashKey = "04DgUStWPtd0aJmiR+L2z5xwbpPr/hmH" //"the-big-and-secret-fash-key-here"//
	CookieBlockKey = "5oR820YTUcjpBFAYlgbte8jEL9o2oHI2" //"lot-secret-of-characters-big-too"//
	CookieExpireSeconds = 86400 //xxx seconds
)

//global variable for single instance
var (
	_cookie *Cookie
	_cookieOnce sync.Once
)

 //face info
 type Cookie struct {
 	secureCookie *securecookie.SecureCookie `secure cookie instance`
 	jwt iface.IJwt
 	expireTime int
 }

//get single instance
func GetCookie() *Cookie {
	_cookieOnce.Do(func() {
		_cookie = NewCookie()
	})
	return _cookie
}

 //construct
func NewCookie() *Cookie {
	//self init
	this := &Cookie{
		expireTime:CookieExpireSeconds,
		jwt: NewJwt(),
	}

	//inter init
	this.interInit()
	return this
}

///////
//api
//////

//set expire time
func (f *Cookie) SetExpireTime(seconds int) {
	f.expireTime = seconds
}

//delete cookie
func (f *Cookie) DelCookie(name,
			domain string,
			c *gin.Context,
		) error {
	//check
	if name == "" || c == nil {
		return errors.New("invalid parameter")
	}
	//destroy cookie
	c.SetCookie(name, "", -1, "/", domain, false, true)
	return nil
}

//get cookie
func (f *Cookie) GetCookie(
			key string,
			c *gin.Context,
		) (map[string]interface{}, error) {
	//check
	if key == "" || c == nil {
		return nil, errors.New("invalid parameter")
	}
	//get original value
	orgVal, err := c.Cookie(key)
	if err != nil {
		return nil, err
	}
	//try decode pass jwt
	jwt, err := f.jwt.Decode(orgVal)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

//set cookie
func (f *Cookie) SetCookie(
			key string,
			val map[string]interface{},
			expireSeconds int,
			domain string,
			c *gin.Context,
		) error {
	//check
	if key == "" || val == nil || c == nil {
		return errors.New("invalid parameter")
	}
	//try encode pass jwt
	encStr, err := f.jwt.Encode(val)
	if err != nil {
		return err
	}
	//set into cookie
	c.SetCookie(key, encStr, expireSeconds, "/", domain, false, true)
	return nil
}

//set jwt
func (f *Cookie) SetJwt(jwt iface.IJwt) bool {
	if jwt == nil {
		return false
	}
	f.jwt = jwt
	return true
}

//////////////
//private func
//////////////

func (f *Cookie) interInit() {
	//init security key
	hashKey := []byte(CookieHashKey)
	blockKey := []byte(CookieBlockKey)

	//init cookie
	f.secureCookie = securecookie.New(hashKey, blockKey)
}