package operations

const (
	Ping = iota
	SetLock
	UnsetLock
	ClaimMaster
	Increment
	Decrement
	Reset
	End
	Nothing
)

type SetLockPayload struct {
	ResourceId  string
	RequesterId string
	TTL         int64
}

type UnsetLockPayload struct {
	ResourceId  string
	RequesterId string
}

type ClaimMasterPayload struct {
	GroupId     string
	RequesterId string
	TTL         int64
}

type IncrementCounterPayload struct {
	CounterId string
}

type DecrementCounterPayload struct {
	CounterId string
}

type ResetCounterPayload struct {
	CounterId string
	Value     int64
}

func FromString(s string) int {
	switch s {
	case "PING":
		return Ping
	case "LOCK":
		return SetLock
	case "UNLOCK":
		return UnsetLock
	case "MASTER":
		return ClaimMaster
	case "INCREMENT":
		return Increment
	case "DECREMENT":
		return Decrement
	case "RESET":
		return Reset
	case "END":
		return End
	}
	// if we don't have a valid command, we do nothing
	return Nothing
}
