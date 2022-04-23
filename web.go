package tinyweb

import (
	"github.com/andyzhou/tinyweb/face"
	"github.com/andyzhou/tinyweb/iface"
	"sync"
)

/*
 * web face info
 */

//global variable
var (
	_web *Web
	_webOnce sync.Once
)

type Web struct {
	app iface.IWebApp
	web iface.IWeb
	captcha iface.ICaptcha
	cookie iface.ICookie
	jwt iface.IJwt
	signature iface.ISignature
	zip iface.IZip
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
	this := &Web{
		app: face.NewWebApp(),
		web: face.NewWeb(),
		captcha: face.NewCaptcha(),
		cookie: face.NewCookie(),
		jwt: face.NewJwt(),
		signature: face.NewSignature(),
		zip: face.NewZip(),
	}
	return this
}

//get relate sub face
func (f *Web) GetZip() iface.IZip {
	return f.zip
}
func (f *Web) GetSignature() iface.ISignature {
	return f.signature
}
func (f *Web) GetJwt() iface.IJwt {
	return f.jwt
}
func (f *Web) GetCookie() iface.ICookie {
	return f.cookie
}
func (f *Web) GetCaptcha() iface.ICaptcha {
	return f.captcha
}
func (f *Web) GetWeb() iface.IWeb {
	return f.web
}
func (f *Web) GetWebApp() iface.IWebApp {
	return f.app
}