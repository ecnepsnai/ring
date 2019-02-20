package ring

import (
	"sync"
	"testing"
)

func TestParallelAdd(t *testing.T) {
	maximum := 15
	ring := New(maximum)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			ring.Add("1")
			i++
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			ring.Add("2")
			i++
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			ring.Add("3")
			i++
		}
	}()

	wg.Wait()

	values := ring.All()
	length := len(values)
	if length != maximum {
		t.Errorf("Ring length is not correct. Expected %d got %d", maximum, length)
	}

	last := ring.Last()

	if last != values[0] {
		t.Error("Incorrect sort order for all ring entries")
	}
}

func TestAll(t *testing.T) {
	ring := New(10)
	ring.Add(1)
	objects := ring.All()
	count := len(objects)
	if count != 1 {
		t.Errorf("Unexpected number of objects returned. Expected %d got %d", 1, count)
	}
}

func TestEmptyAll(t *testing.T) {
	ring := New(15)
	all := ring.All()
	length := len(all)
	if length != 0 {
		t.Error("Non-empty array returned for empty ring")
	}
}

func TestEmptyLast(t *testing.T) {
	ring := New(15)
	last := ring.Last()
	if last != nil {
		t.Error("Non-nil last value returne for empty ring")
	}
}
