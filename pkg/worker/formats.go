package worker

import (
	"fmt"
	"github.com/google/uuid"
)

type Operation struct {
	Type int8
	Body interface{}
}

type Response struct {
	Id  uuid.UUID // ids to match requests and responses in logging
	Err error     // if error is not nil, we return error message
	Res []byte    // response message
}

type Request struct {
	Id    uuid.UUID     // ids to match requests and responses in logging
	Op    *Operation    // message to process
	ResCh chan Response // channel to send response to
}

// ----
// helper functions

// workerError returns a standard error message
func workerError(id uuid.UUID, err error) *Response {
	fmt.Println("Worker error", err)
	return &Response{Id: id, Err: err, Res: []byte{}}
}
