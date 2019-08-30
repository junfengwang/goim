package service

import (
	"net"
	"time"
	"context"
)

type Conn struct {
	conn net.Conn
	sendMsgCh chan []byte
	receiveCh chan *Message
	doneCh    chan error
	hbTimeout time.Duration
}

func NewConn(conn net.Conn, hbTimeout time.Duration) *Conn {
	c := &Conn{
		conn:conn,
		hbTimeout:hbTimeout,
		sendMsgCh:make(chan []byte, 10),
		receiveCh:make(chan *Message, 10),
		doneCh:make(chan error),
	}

	return c
}


//发送消息
func (c *Conn) WriteHandler(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			return
		case msg := <- c.sendMsgCh:
			if msg == nil {
				continue
			}

			_, err := c.conn.Write(msg)
			if err != nil {
				c.doneCh <- err
				return
			}

		}
	}
}

func (c *Conn) ReadHandler(ctx context.Context) {
	for {
		select {
			case ctx.Done():
				return
			default:

		}
	}
}