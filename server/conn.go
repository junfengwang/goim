package server

import "net"

type Conn struct {
	uid string
	name string
	conn net.Conn
}
