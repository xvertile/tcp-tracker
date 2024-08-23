package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"xvertile/tcp-tracker/tracker"
)

func main() {
	conn, err := net.Dial("tcp", "icanhazip.com:80")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	maxBytes := int64(1024)
	trackedConn := tracker.CreateCountingConn(conn, maxBytes)
	request := "GET / HTTP/1.1\r\nHost: icanhazip.com\r\nConnection: close\r\n\r\n"
	_, err = trackedConn.Write([]byte(request))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	response := make([]byte, 4096)
	for {
		n, err := trackedConn.Read(response)
		if err != nil && err != io.EOF {
			log.Fatalf("Failed to read response: %v", err)
		}
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Print(string(response[:n]))
	}
	log.Printf("Total bytes used: %d\n", trackedConn.BytesRead)
}
