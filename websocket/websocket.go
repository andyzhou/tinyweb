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
	wsRootUri string
	gin *gin.Engine //parent reference
	upgrade websocket.Upgrader
	connManager IConnManager
	coder ICoder
}

//construct
func NewWebSocket() *WebSocket {
	this := &WebSocket{
		connManager: NewManager(),
		coder: NewCoder(),
	}
	return this
}

//set gin engine
func (f *WebSocket) SetGin(gin *gin.Engine) {
	f.gin = gin
}

//register web socket request uri
func (f *WebSocket) RegisterWs(rootUri string) {
	////check
	//
	////set gin and websocket root uri
	//wsRootUri := define.WebSocketRoot
	//if rootUri != nil && len(rootUri) > 0 {
	//	wsRootUri = rootUri[0]
	//}
	//f.wsRootUri = wsRootUri
	//f.gin = gin
	f.gin.GET(rootUri, f.processConn)
}

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
	session := c.Query(define.QueryParaOfSession)
	//contentType := c.Query(define.QueryParaOfContentType)

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
	wsConn, err := f.connManager.Accept(session, conn)
	if err != nil {
		err = f.connManager.CloseWithMessage(conn, define.MessageForNormalClosed)
		if err != nil {
			log.Printf("WebSocketServer:processRequest, err:%v", err.Error())
		}
		conn.Close()
		return
	}

	//loop read data
	for {
		_, _, err := wsConn.Read()
		if err != nil {
			// handle error
			if err == io.EOF {
				log.Printf("ws EOF need close!")
				return
			}
			log.Printf("ws err need close, err:%v", err.Error())
			return
		}
	}
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
	if f.gin == nil {
		panic("WebSocket:interInit, gin instance is nil")
		return
	}
}