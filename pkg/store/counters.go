package store

type Counter struct {
	Value int64
}

// counters is a map of resource_id to Counter
var counters = map[string]*Counter{}

func Increment(resourceId string,) Counter {
	counter, ok := counters[resourceId]

	if !ok {
		r := &Counter{Value: 1}
		counters[resourceId] = r
	} else {
		counter.Value++
	}

	// return a copy of the counter
	return *counter
}

func Decrement(resourceId string) (Counter, bool) {
	counter, ok := counters[resourceId]

	if !ok {
		r := &Counter{Value: 0}
		counters[resourceId] = r
		return *r, true
	} else {
		if counter.Value == 0 {
			return *counter, false
		} else {
			counter.Value--
			return *counter, true
		}
	}
}

func Reset(resourceId string, value int64) (Counter, bool) {
	counter, ok := counters[resourceId]

	if !ok {
		r := &Counter{Value: value}
		counters[resourceId] = r
		return *r, true
	} else {
		counter.Value = value
		return *counter, true
	}
}
