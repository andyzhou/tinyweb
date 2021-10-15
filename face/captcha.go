package face

import (
	"fmt"
	"github.com/afocus/captcha"
	"github.com/kataras/iris/v12"
	"image/color"
	"image/png"
	"os"
)
/*
 * face for captcha
 * @author <AndyZhou>
 * @mail <diudiu8848@163.com>
 */

//inter macro define
const (
	CaptchaWidth = 128
	CaptchaHeight = 48
	CaptchaNumSize = 5
)

//face info
type Captcha struct {
	cookieName string
	cookieExpire int
	numSize int
	cap *captcha.Captcha
	cookie *Cookie
}

//construct
func NewCaptcha() *Captcha {
	//self init
	this := &Captcha{
		cap: captcha.New(),
		numSize: CaptchaNumSize,
	}
	//inter init
	this.interInit()
	return this
}

//set image size
func (f *Captcha) SetSize(width, height int) bool {
	if width <= 0 || height <= 0 {
		return false
	}
	//set captcha size
	f.cap.SetSize(width, height)
	return true
}

//set font path
func (f *Captcha) SetFontPath(path string) bool {
	if path == "" {
		return false
	}
	fontFilePath := fmt.Sprintf("%s/comic.ttf", path)
	f.cap.SetFont(fontFilePath)
	return true
}

//captcha num, default 5
func (f *Captcha) SetCaptchaNum(num int) bool {
	if num <= 0 {
		return false
	}
	f.numSize = num
	return false
}

//set cookie config
func (f *Captcha) SetCookiePara(name string, expire int) bool {
	if name == "" || expire < 0 {
		return false
	}
	f.cookieName = name
	f.cookieExpire = expire
	return true
}

//get cookie value
func (f *Captcha) GetCookie(ctx iris.Context) string {
	return f.cookie.GetCookie(f.cookieName, ctx)
}

//gen image
func (f *Captcha) GenImg(ctx iris.Context) {
	//create captcha image
	img, str := f.cap.Create(f.numSize, captcha.NUM)

	//set cookie
	f.cookie.SetCookie(f.cookieName, str, f.cookieExpire, ctx)

	//get writer
	writer := ctx.ResponseWriter()
	writer.Header().Set("Content-Type", "image/png")

	//output png image
	png.Encode(writer, img)
}

///////////////
//private func
///////////////

//inter init
func (f *Captcha) interInit() {
	//init cookie
	f.cookie = NewCookie()

	//get current pwd
	pwd, _ := os.Getwd()

	//set default values
	f.SetCaptchaNum(CaptchaNumSize)
	f.SetFontPath(pwd)
	f.SetSize(CaptchaWidth, CaptchaHeight)

	//set disturbance
	f.cap.SetDisturbance(captcha.MEDIUM)

	//set front and back color
	f.cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	f.cap.SetBkgColor(
						color.RGBA{255, 0, 0, 255},
						color.RGBA{0, 0, 255, 255},
						color.RGBA{0, 153, 0, 255},
					)
}
