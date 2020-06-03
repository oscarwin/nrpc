package transport

import (
	"net"
	"time"
)

type NSocket struct {
	conn    net.Conn
	timeout time.Duration
}

func (s *NSocket) Open() error {
	return nil
}

func (s *NSocket) IsOpen() bool {
	return s.conn != nil
}

func (s *NSocket) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s *NSocket) Write(p []byte) (n int, err error) {
	return 0, nil
}
