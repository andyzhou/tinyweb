package websocket

import (
	"github.com/google/uuid"
	"time"
)

//face info
type Session struct {
	expireSec time.Duration
}

//construct
func NewSession() *Session {
	this := &Session{}
	return this
}

//set expire second
func (s *Session) SetExpire(exp time.Duration) {
	s.expireSec = exp
}

//create session
func (s *Session) CreateSession(userId string) (string, error) {
	session := s.generateUUId()
	return session, nil
}


//generate uuid
func (s *Session) generateUUId() string {
	return uuid.New().String()
}