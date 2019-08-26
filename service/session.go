package service

type Session struct {
	conn Conn
	uid int32
	config []interface{}
}
