package service

import (
	"sync"
	"net"
	"time"
	"fmt"
	"context"
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

func (s *Socketserver) RegisterOnMessage(handler func (*Session, *Message) ) {
	s.onMessage = handler
}

func (s *Socketserver) ServStar() {
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		cancel()

		s.listener.Close()
	}()

	go s.StartAccept(ctx)

	for {
		select {
			case <-s.stopCh:
				return
		}
	}
}

func (s *Socketserver) StartAccept(ctx context.Context) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}
	}
}