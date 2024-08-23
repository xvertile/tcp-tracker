package tracker

import (
	"net"
)

func (c *CountingConn) Read(b []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	n, err := c.Conn.Read(b)
	if n > 0 {
		c.BytesRead += int64(n)
		if c.BytesRead > c.MaxBytes {
			c.Conn.Close()
			return n, net.ErrClosed
		}
	}

	return n, err
}

func (c *CountingConn) Write(b []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	n, err := c.Conn.Write(b)
	if n > 0 {
		c.BytesRead += int64(n)
		if c.BytesRead > c.MaxBytes {
			c.Conn.Close()
			return n, net.ErrClosed
		}
	}

	return n, err
}

func CreateCountingConn(conn net.Conn, maxBytes int64) *CountingConn {
	return &CountingConn{Conn: conn, MaxBytes: maxBytes}
}
