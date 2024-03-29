package websocket

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

/*
 * websocket manager
 */

//manager info
type Manager struct {
	connMap sync.Map //session -> wsConn
}

//construct
func NewManager() *Manager {
	this := &Manager{
		connMap: sync.Map{},
	}
	return this
}

//get conn by session
func (f *Manager) GetConnBySession(session string) *websocket.Conn {
	if session == "" {
		return nil
	}
	v, ok := f.connMap.Load(session)
	if !ok || v == nil {
		return nil
	}
	conn, ok := v.(*websocket.Conn)
	if !ok {
		return nil
	}
	return conn
}

//accept websocket connect
func (f *Manager) Accept(conn *websocket.Conn, session, userId string) (IWSConn, error) {
	//check
	if session == "" || conn == nil {
		return nil, errors.New("invalid parameter")
	}

	//check session, todo..

	//init new connect
	wsConn := NewWSConn(conn)
	f.connMap.Store(session, wsConn)

	return wsConn, nil
}

//close conn with message
func (f *Manager) CloseWithMessage(conn *websocket.Conn, message string) error {
	msg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, message)
	err := conn.WriteMessage(websocket.CloseMessage, msg)
	if err != nil {
		return err
	}
	return conn.Close()
}

//close conn
func (f *Manager) CloseConn(session string) error {
	//check
	if session == "" {
		return errors.New("invalid parameter")
	}
	//load and update
	v, ok := f.connMap.Load(session)
	if !ok || v == nil {
		return errors.New("no record for current session")
	}
	wsConn, ok := v.(*WSConn)
	if !ok || wsConn == nil {
		return errors.New("invalid data format")
	}
	//begin close and clear
	err := wsConn.Close()
	if err != nil {
		return err
	}
	f.connMap.Delete(session)
	return nil
}