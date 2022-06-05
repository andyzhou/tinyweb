package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

/*
 * websocket relate interface define
 */

//interface of websocket
type IWebSocket interface {
	RegisterWs(reqUrl string, cbForRead func(msgType int, data []byte) ([]byte, error))
	SetCheckCB(cbForCheck func(c *gin.Context) bool)
	SetConnCB(cbForConn func(c *gin.Context))
	SetKeyPara(session, userId string)
	SetGin(gin *gin.Engine)
}

//interface of manager
type IConnManager interface {
	GetConnBySession(session string) *websocket.Conn
	Accept(conn *websocket.Conn, session, userId string) (IWSConn, error)
	CloseWithMessage(conn *websocket.Conn, message string) error
	CloseConn(session string) error
}

//interface of connect
type IWSConn interface {
	Write(messageType int, data []byte) error
	Read() (int, []byte, error)
	CloseWithMessage(message string) error
	Close() error
}

//interface for message en/decoder
type ICoder interface {
	Marshal(contentType string, content proto.Message) ([]byte, error)
	Unmarshal(contentType string, content []byte, req proto.Message) error
}