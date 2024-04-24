package listener

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mastermind/pkg/failures"
	"mastermind/pkg/parser"
	"mastermind/pkg/worker"
	"net"
	"strings"
	"time"
)

const maxSizeMessage = 5000 // 5000 char maximum message

func processConnection(conn net.Conn, w chan worker.Request) {
	var err error
	defer wrap(conn)
	fmt.Println("New connection from", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	message := ""

	// set read deadline
	// TODO make this configurable via a global configuration
	if err = conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		failures.ToStderr(failures.DEADLINE_ERROR, err)
		return
	}

	for {
		// read from connection into buffer
		// add to the message string bit by bit
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				failures.ToStderr(failures.CONN_READ_ERROR, err)
				write(conn, []byte(failures.CONN_READ_ERROR))
				return
			}
			break
		}

		message += string(buf[:n])
		buf = make([]byte, 1024) // clear buffer

		if strings.Contains(message, "EOF") || n == 0 || len(message) > maxSizeMessage {
			break
		}
	}

	// TODO make logs debug level ???
	fmt.Println("Received data:", message, "from", conn.RemoteAddr().String())

	op, err := parser.Parse(message)
	if err != nil {
		failures.ToStderr(failures.PARSE_ERROR, err)
		write(conn, []byte(failures.PARSE_ERROR))
		return
	}

	r := make(chan worker.Response)
	defer close(r)

	w <- worker.Request{
		Id:    uuid.New(),
		Op:    op,
		ResCh: r,
	}

	// read from the worker channel with a timeout
	// TODO make this configurable via a global configuration
	select {
	case <-time.After(3 * time.Second): // three seconds feels neat
		failures.ToStderr(failures.WORKER_TIMEOUT, nil)
		write(conn, []byte(failures.WORKER_TIMEOUT))
		return
	case res := <-r:
		write(conn, res.Res)
	}
}
