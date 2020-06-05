package transport

import (
	"fmt"
	"net"
	"time"
)

type NSocketTCP struct {
	conn    net.Conn
	timeout time.Duration
}

// create a NSocketTCP from an existing net.Conn
func NewNSocketWithConn(conn net.Conn, timeout time.Duration) *NSocketTCP {
	return &NSocketTCP{
		conn:    conn,
		timeout: timeout,
	}
}

// create a NSocketTCP from an existing net.Conn and a timeout
func NewNSocketWithConnTimeout(conn net.Conn, timeout time.Duration) *NSocketTCP {
	return &NSocketTCP{
		conn:    conn,
		timeout: timeout,
	}
}

// set the socket timeout
func (s *NSocketTCP) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *NSocketTCP) Open() error {
	return nil
}

// check the connection is open
func (s *NSocketTCP) IsOpen() bool {
	return s.conn != nil
}

func (s *NSocketTCP) Read(buf []byte) (n int, err error) {
	if !s.IsOpen() {
		return 0, fmt.Errorf("connection not open")
	}
	s.pushDeadline(true, false)
	return s.conn.Read(buf)
}

func (s *NSocketTCP) Write(p []byte) (n int, err error) {
	if !s.IsOpen() {
		return 0, fmt.Errorf("connection not open")
	}
	s.pushDeadline(false, true)
	return s.conn.Write(p)
}

func (s *NSocketTCP) Close() error {
	if s.conn != nil {
		if err := s.conn.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (s *NSocketTCP) pushDeadline(read, write bool) {
	var t time.Time
	if s.timeout > 0 {
		t = time.Now().Add(s.timeout)
	}
	if read && write {
		s.conn.SetDeadline(t)
	} else if read {
		s.conn.SetReadDeadline(t)
	} else if write {
		s.conn.SetWriteDeadline(t)
	}
}
