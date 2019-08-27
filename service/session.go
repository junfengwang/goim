package service

type Session struct {
	conn Conn
	uid int32
	config []interface{}
}

func NewSession(c Conn) (*Session) {

	session := &Session{
		conn:c,
	}

	return session
}
