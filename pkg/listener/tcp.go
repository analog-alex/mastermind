package listener

import (
	"mastermind/pkg/failures"
	"mastermind/pkg/worker"
	"net"
)

// store worker channel as a folder global
var wc chan worker.Request

func init() {
	wc = worker.New()
}

// TPCListener create a tcp listener on given port
func TPCListener(port string) {
	// set up tcp listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		failures.ToStderr(failures.SERVER_BOOT_ERROR, err)
		return
	}

	// close listener on exit
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			failures.ToStderr(failures.LISTENER_CLOSING_ERROR, err)
		}
	}(listener)

	for {
		// accept connections
		conn, err := listener.Accept()
		if err != nil {
			failures.ToStderr(failures.ACCEPT_ERROR, err)
			continue
		}

		// process connection in a new goroutine
		go processConnection(conn, wc)
	}
}
