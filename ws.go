package tinyweb

import (
	"errors"
	"github.com/andyzhou/tinyweb/websocket"
	"github.com/gin-gonic/gin"
	"sync"
)

/*
 * websocket face info
 */

//global variable
var (
	_websocket *WebSocket
	_websocketOnce sync.Once
)

//face info
type WebSocket struct {
	ws websocket.IWebSocket
}

//construct
func NewWebSocket(g ...*gin.Engine) *WebSocket {
	this := &WebSocket{
		ws: websocket.NewWebSocket(g...),
	}
	return this
}

//send message

//register ws url and cb
func (f *WebSocket) RegisterWS(
						wsUrl string,
						cbForRead func(int, []byte) error,
					) error {
	//check
	if wsUrl == "" || cbForRead == nil {
		return errors.New("invalid parameter")
	}
	f.ws.RegisterWs(wsUrl, cbForRead)
	return nil
}

//set cb for check connect
func (f *WebSocket) SetCBForCheckConn(cb func(c *gin.Context) bool) {
	f.ws.SetCheckCB(cb)
}

//set cb for connect
func (f *WebSocket) SetCBForConn(cb func(c *gin.Context)) {
	f.ws.SetConnCB(cb)
}

//set key param
func (f *WebSocket) SetKeyPara(session, userId string) {
	f.ws.SetKeyPara(session, userId)
}

//set gin engine
func (f *WebSocket) SetGin(g *gin.Engine) {
	f.ws.SetGin(g)
}
