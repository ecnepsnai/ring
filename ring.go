// Package ring is a threadsafe rotating slice. You define a maximum number of items
// and for each added item, it will fill the slice up to that max, then begin
// to wrap around the start of the slice.
package ring

import (
	"sync"
)

// Ring describes a single ring instance
type Ring struct {
	lock       *sync.RWMutex
	currentIdx int
	maximumIdx int
	values     []interface{}
}

// New will create a new ring with the specified maximum number of items
func New(Maximum int) *Ring {
	return &Ring{
		lock:       &sync.RWMutex{},
		currentIdx: -1,
		maximumIdx: Maximum - 1,
		values:     make([]interface{}, Maximum),
	}
}

// Add will add a new object to the ring
func (r *Ring) Add(object interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()

	var nextIdx int
	if r.currentIdx >= r.maximumIdx {
		nextIdx = 0
	} else {
		nextIdx = r.currentIdx + 1
	}

	r.values[nextIdx] = object
	r.currentIdx = nextIdx
}

// All will return a copy of all objects in the ring sorted by when they were added
func (r *Ring) All() []interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if r.currentIdx == -1 {
		return []interface{}{}
	}

	var length int
	for i, obj := range r.values {
		if obj != nil {
			length = i + 1
		}
	}

	var copy = make([]interface{}, length)
	idx := r.currentIdx
	got := 0
	for got < length {
		copy[idx] = r.values[idx]
		if idx == 0 {
			idx = r.maximumIdx
		} else {
			idx--
		}
		got++
	}

	return copy
}

// Last will return the last inserted object in the ring, or nil if the ring is empty
func (r *Ring) Last() interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if r.currentIdx == -1 {
		return nil
	}

	return r.values[r.currentIdx]
}

// Truncate will remove all objects from the ring
func (r *Ring) Truncate() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.currentIdx = -1
	r.values = make([]interface{}, r.maximumIdx+1)
}
