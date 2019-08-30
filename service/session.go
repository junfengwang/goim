package service

import "github.com/satori/go.uuid"

type Session struct {
	conn *Conn
	uid string
	config []interface{}
}

func NewSession(c *Conn) (*Session) {
	uid, _ := uuid.NewV4()
	session := &Session{
		conn:c,
		uid:uid.String(),
	}

	return session
}
