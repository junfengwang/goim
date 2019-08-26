package service

import (
	"net"
	"time"
)

type Conn struct {
	conn net.Conn
	sendMsgCh chan []byte
	receiveCh chan *Message
	doneCh    chan error
	hbTimeout time.Duration
}
