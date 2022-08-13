package websocket

import (
	"github.com/andyzhou/tinyweb/define"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"runtime/debug"
)

/*
 * face of websocket
 * - http base on github.com/gin-gonic/gin
 * - ws base one github.com/gorilla/websocket
 */

//upgrade http connect to ws connect
var upGrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//face info
type WebSocket struct {
	//cb for read data
	cbForRead func(int, []byte) ([]byte, error)

	//cb for basic check
	cbForCheck func(c *gin.Context) bool

	//cb for new connect
	cbForConnect func(c *gin.Context)

	//key para
	sessionKey string
	userIdKey string

	//inter reference
	server *gin.Engine //parent reference
	upgrade websocket.Upgrader
	connManager IConnManager
	coder ICoder
	wsRootUri string
}

//construct
func NewWebSocket(g ...*gin.Engine) *WebSocket {
	var (
		s *gin.Engine
	)
	//init default gin engine
	if g != nil && len(g) > 0 {
		s = g[0]
	}else{
		s = gin.Default()
	}
	//self init
	this := &WebSocket{
		server: s,
		connManager: NewManager(),
		coder: NewCoder(),
		sessionKey: define.QueryParaOfSession,
		userIdKey: define.QueryParaOfUserId,
	}
	//inter init
	this.interInit()
	return this
}

//set gin engine
func (f *WebSocket) SetGin(gin *gin.Engine) {
	f.server = gin
}

//set key param
func (f *WebSocket) SetKeyPara(session, userId string) {
	if session != "" {
		f.sessionKey = session
	}
	if userId != "" {
		f.userIdKey = userId
	}
}

//set cb for new connect
func (f *WebSocket) SetConnCB(cbForConn func(c *gin.Context)) {
	if f.cbForConnect != nil {
		return
	}
	f.cbForConnect = cbForConn
}

//set cb for check
func (f *WebSocket) SetCheckCB(cbForCheck func(c *gin.Context) bool) {
	if f.cbForCheck != nil {
		return
	}
	f.cbForCheck = cbForCheck
}

//register web socket request uri
func (f *WebSocket) RegisterWs(
						reqUrl string,
						cbForRead func(int, []byte) ([]byte, error),
					){
	if f.cbForRead != nil {
		return
	}
	f.cbForRead = cbForRead
	f.server.GET(reqUrl, f.processConn)
}

///////////////
//private func
///////////////

//process web socket connect
func (f *WebSocket) processConn(c *gin.Context) {
	//defer
	defer func() {
		if err := recover(); err != nil {
			log.Printf("WebSocketServer:processRequest panic, err:%v, stack:%v",
				err, string(debug.Stack()))
		}
	}()

	//get key data
	request := c.Request
	writer := c.Writer

	//get key param
	session := c.Query(f.sessionKey)
	userId := c.Query(f.userIdKey)

	//cb for connect
	if f.cbForConnect != nil {
		f.cbForConnect(c)
	}

	//cb for check
	if f.cbForCheck != nil {
		if !f.cbForCheck(c) {
			return
		}
	}

	//setup net base data
	//netBase := &NetBase{
	//	ContentType: contentType,
	//	ClientIP: c.ClientIP(), //get client id
	//}

	//upgrade http connect to ws connect
	conn, err := f.upgrade.Upgrade(writer, request, nil)
	if err != nil {
		//500 error
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//accept new connect
	wsConn, err := f.connManager.Accept(conn, session, userId)
	if err != nil {
		err = f.connManager.CloseWithMessage(conn, define.MessageForNormalClosed)
		if err != nil {
			log.Printf("WebSocketServer:processRequest, err:%v", err.Error())
		}
		conn.Close()
		return
	}

	//loop read data
	go f.readConn(wsConn, session, userId)
}

//read process for per connect
func (f *WebSocket) readConn(wsConn IWSConn, session, userId string) error {
	//defer
	defer func() {
		if err := recover(); err != nil {
			log.Printf("read connect panic, err:%v", err)
		}
	}()

	//loop read data
	for {
		//messageType, data, error
		msgType, data, err := wsConn.Read()
		if err != nil {
			// handle error
			if err == io.EOF {
				log.Printf("ws EOF need close!")
				return nil
			}
			log.Printf("ws err need close, err:%v", err.Error())
			return err
		}
		//call cb for read
		if f.cbForRead != nil {
			resp, err := f.cbForRead(msgType, data)
			if err != nil {
				log.Printf("ws read resp failed, err:%v", err)
				continue
			}
			//send response
			err = wsConn.Write(msgType, resp)
			if err != nil {
				log.Printf("ws write resp failed, err:%v", err)
			}
		}
	}
	return nil
}

//inter init
func (f *WebSocket) interInit() {
	//init websocket upgrade
	f.upgrade = websocket.Upgrader{
		ReadBufferSize: define.WebSocketBufferSize,
		WriteBufferSize: define.WebSocketBufferSize,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//setup websocket
	if f.server == nil {
		panic("WebSocket:interInit, gin instance is nil")
		return
	}
}