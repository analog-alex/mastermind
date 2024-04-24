package store

import (
	"time"
)

type Resource struct {
	RequesterId string
	StoredAt    time.Time
	TTL         time.Duration
	Enforced    bool
	Order       int
}

// locks is a map of resource_id to Resource
var locks = map[string]*Resource{}

// SetLock sets a lock for a given group id and requester id
// if the lock is already set, it will return the current lock and a boolean
// indicating if the requester is the lock owner
func SetLock(resourceId string, requesterId string, ttl int64) (Resource, bool) {
	resource, ok := locks[resourceId]

	// if we no lock or the lock is expired we accept the claim
	if !ok || time.Since(resource.StoredAt) > resource.TTL {
		r := &Resource{
			RequesterId: requesterId,
			StoredAt:    time.Now(),
			TTL:         time.Second * time.Duration(ttl),
			Enforced:    false,
			Order:       1,
		}
		locks[resourceId] = r
		return *r, true
	}

	// if the lock is not expired, we return the current lock and a boolean indicating if the requester is the owner
	return *resource, resource.RequesterId == requesterId
}

// UnsetLock unsets a lock for a given group id and requester id
// if the lock is set and the requester is the owner, it will be unset, and we return true
// if the lock is not set, we return true as well (resource is "free")
// if the lock is set and the requester is not the owner, we return false
func UnsetLock(resourceId string, requesterId string) bool {
	resource, ok := locks[resourceId]

	if ok {
		if resource.RequesterId == requesterId {
			delete(locks, resourceId)
			return true
		}
		return false
	}
	return true
}

// GetLock returns the lock for a given group id, if it exists
func GetLock(resourceId string) (Resource, bool) {
	resource, ok := locks[resourceId]
	return *resource, ok
}

// IsLocked returns true if the resource is locked
// false otherwise or if the resource does not exist
func IsLocked(resourceId string) bool {
	_, ok := locks[resourceId]
	return ok
}
