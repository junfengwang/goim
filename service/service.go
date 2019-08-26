package service

import (
	"sync"
	"net"
	"time"
	"fmt"
)

type Socketserver struct {
	sessions        *sync.Map
	listener        net.Listener
	heatbeatTimeout time.Duration
	status          int
	stopCh          chan error
	onMessage       func(*Session, *Message)
	onConnect       func(*Session)
	onDisConnect    func(*Session)
}


func NewSocketserver(laddr string) (*Socketserver, error) {

	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		fmt.Println("listen error", err)

		return nil, err
	}

	s := &Socketserver{
		sessions: &sync.Map{},
		listener:listener,
		heatbeatTimeout:0 * time.Second,
	}

	return s, nil
}

func (s *Socketserver) RegisterOnMessage(handler func (*session, *Message) ) {

}

func (s *Socketserver) StartAccept() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}
	}
}