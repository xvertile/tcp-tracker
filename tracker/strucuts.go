package tracker

import (
	"net"
	"sync"
)

type CountingConn struct {
	net.Conn
	mu        sync.Mutex
	BytesRead int64
	MaxBytes  int64
}
