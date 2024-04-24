package failures

const (
	SERVER_BOOT_ERROR      = "error booting up server"
	LISTENER_CLOSING_ERROR = "error closing TCP listener"
	ACCEPT_ERROR           = "failure when accepting a connection"
	DEADLINE_ERROR         = "error setting read deadline"
	CONN_READ_ERROR        = "error reading from connection"
	PARSE_ERROR            = "error parsing request"
	WORKER_TIMEOUT         = "timeout waiting for response from worker"
	INVALID_COMMAND        = "invalid command"
)
