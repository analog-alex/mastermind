package parser

import (
	"fmt"
	"mastermind/pkg/failures"
	"mastermind/pkg/worker"
	"mastermind/pkg/worker/operations"
	"strings"
)

func Parse(cmd string) (*worker.Operation, error) {
	s := strings.Split(cmd, " ")
	if len(s) != 4 {
		return &worker.Operation{}, fmt.Errorf(failures.INVALID_COMMAND)
	}

	o := operations.FromString(strings.ToUpper(s[0]))

	switch o {

	case operations.Ping:
		return &worker.Operation{Type: operations.Ping, Body: 1}, nil

	case operations.SetLock:
		body := &operations.SetLockPayload{ResourceId: s[1], RequesterId: s[2], TTL: 10}
		return &worker.Operation{Type: operations.SetLock, Body: body}, nil

	case operations.UnsetLock:
		body := &operations.UnsetLockPayload{ResourceId: s[1], RequesterId: s[2]}
		return &worker.Operation{Type: operations.UnsetLock, Body: body}, nil

	case operations.ClaimMaster:
		body := &operations.ClaimMasterPayload{GroupId: s[1], RequesterId: s[2], TTL: 10}
		return &worker.Operation{Type: operations.ClaimMaster, Body: body}, nil

	case operations.Increment:
		body := &operations.IncrementCounterPayload{CounterId: s[1]}
		return &worker.Operation{Type: operations.Increment, Body: body}, nil

	case operations.Decrement:
		body := &operations.DecrementCounterPayload{CounterId: s[1]}
		return &worker.Operation{Type: operations.Decrement, Body: body}, nil

	case operations.Reset:
		body := &operations.ResetCounterPayload{CounterId: s[1]}
		return &worker.Operation{Type: operations.Reset, Body: body}, nil

	default:
		return &worker.Operation{}, fmt.Errorf(failures.INVALID_COMMAND)

	}
}
