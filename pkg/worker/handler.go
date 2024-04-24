package worker

import (
	"fmt"
	"mastermind/pkg/store"
	"mastermind/pkg/worker/operations"
)

// handeMessage handles messages from the worker queue
// we want to process messages sequentially
func handeMessage(req Request) (Response, error, bool) {
	switch req.Op.Type {

	// ----------------------
	// ping case
	case operations.Ping:
		return Response{Id: req.Id, Err: nil, Res: []byte("pong")}, nil, false

	// ----------------------
	// is master case
	case operations.SetLock:
		// cast body to SetLockPayload
		body, ok := req.Op.Body.(*operations.SetLockPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		_, is := store.SetLock(body.ResourceId, body.RequesterId, body.TTL)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%t", is))}, nil, false

	// ----------------------
	// unset lock case
	case operations.UnsetLock:
		// cast body to UnsetLockPayload
		body, ok := req.Op.Body.(*operations.UnsetLockPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		is := store.UnsetLock(body.ResourceId, body.RequesterId)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%t", is))}, nil, false

	// --------
	// claim master case
	case operations.ClaimMaster:
		// cast body to ClaimMasterPayload
		body, ok := req.Op.Body.(*operations.ClaimMasterPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		r, is := store.SetLock(body.GroupId, body.RequesterId, body.TTL)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%v - %t", r, is))}, nil, false

	// ----------------------
	// increment counter case
	case operations.Increment:
		// cast body to IncrementCounterPayload
		body, ok := req.Op.Body.(*operations.IncrementCounterPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		r := store.Increment(body.CounterId)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%v", r.Value))}, nil, false

	// ----------------------
	//decrement counter case
	case operations.Decrement:
		// cast body to DecrementCounterPayload
		body, ok := req.Op.Body.(*operations.DecrementCounterPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		r, is := store.Decrement(body.CounterId)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%v - %t", r.Value, is))}, nil, false

	// ----------------------
	// reset counter case
	case operations.Reset:
		// cast body to ResetCounterPayload
		body, ok := req.Op.Body.(*operations.ResetCounterPayload)

		if !ok {
			return Response{}, fmt.Errorf("invalid body"), false
		}

		r, is := store.Reset(body.CounterId, body.Value)
		return Response{Id: req.Id, Err: nil, Res: []byte(fmt.Sprintf("%v - %t", r.Value, is))}, nil, false

	// ----------------------
	// end case -- worker shutdown
	case operations.End:
		return Response{Id: req.Id, Err: nil, Res: []byte("OK")}, nil, true

	// ----------------------
	// default case -- should never happen if our enums are working properly
	default:
		return Response{Id: req.Id, Err: fmt.Errorf("unknown operation"), Res: []byte{}}, nil, false
	}
}
