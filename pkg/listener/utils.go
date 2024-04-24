package listener

import (
	"fmt"
	"mastermind/pkg/failures"
	"net"
)

func write(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("Error writing to connection", err)
	}
}

// wrap closes a connection
func wrap(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		failures.ToStderr("error closing connection", err)
		return // don't panic, simply return, contain error to this goroutine
	}
}
