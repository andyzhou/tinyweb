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
type IWebsocket interface {
	RegisterWs(gin *gin.Engine, rootUri ...string)
	ProcessConn(c *gin.Context)
	GetRootUri() string
}

//interface of manager
type IConnManager interface {
	GetConnBySession(session string) *websocket.Conn
	Accept(session string, conn *websocket.Conn) (IWSConn, error)
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