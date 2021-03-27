package face

import (
	"fmt"
	"github.com/afocus/captcha"
	"github.com/kataras/iris/v12"
	"image/color"
	"image/png"
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
	fontPath string
	cookieName string
	cookieExpire int
	cap *captcha.Captcha
	cookie *Cookie
}

//construct
func NewCaptcha() *Captcha {
	//self init
	this := &Captcha{
		cap: captcha.New(),
	}
	//inter init
	this.interInit()
	return this
}

//set cookie config
func (f *Captcha) SetCookie(name string, expire int) bool {
	if name == "" || expire < 0 {
		return false
	}
	f.cookieName = name
	f.cookieExpire = expire
	return true
}

//gen image
func (f *Captcha) GenImg(ctx iris.Context) {
	//create captcha image
	img, str := f.cap.Create(CaptchaNumSize, captcha.NUM)

	//set cookie
	f.cookie.SetCookie(f.cookieName, str, f.cookieExpire)

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

	//set font
	fontFilePath := fmt.Sprintf("%s/comic.ttf", f.fontPath)
	f.cap.SetFont(fontFilePath)

	//set captcha size
	f.cap.SetSize(CaptchaWidth, CaptchaHeight)

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
