package ring_test

import (
	"sync"
	"testing"

	"github.com/ecnepsnai/ring"
)

func TestParallelAdd(t *testing.T) {
	t.Parallel()

	maximum := 15
	slice := ring.New(maximum)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			slice.Add("1")
			i++
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			slice.Add("2")
			i++
		}
	}()
	go func() {
		defer wg.Done()
		i := 0
		for i < maximum+1 {
			slice.Add("3")
			i++
		}
	}()

	wg.Wait()

	values := slice.All()
	length := len(values)
	if length != maximum {
		t.Errorf("Ring length is not correct. Expected %d got %d", maximum, length)
	}

	last := slice.Last()

	if last != values[0] {
		t.Error("Incorrect sort order for all ring entries")
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	ring := ring.New(10)
	ring.Add(1)
	objects := ring.All()
	count := len(objects)
	if count != 1 {
		t.Errorf("Unexpected number of objects returned. Expected %d got %d", 1, count)
	}
}

func TestEmptyAll(t *testing.T) {
	t.Parallel()

	slice := ring.New(15)
	all := slice.All()
	length := len(all)
	if length != 0 {
		t.Error("Non-empty array returned for empty ring")
	}
}

func TestEmptyLast(t *testing.T) {
	t.Parallel()

	slice := ring.New(15)
	last := slice.Last()
	if last != nil {
		t.Error("Non-nil last value returned for empty ring")
	}
}

func TestTruncate(t *testing.T) {
	t.Parallel()

	slice := ring.New(5)
	i := 0
	for i < 10 {
		slice.Add(i)
		i++
	}

	length := len(slice.All())
	if length != 5 {
		t.Errorf("Unexpected length returned. Expected 5 got %d", length)
	}

	slice.Truncate()
	length = len(slice.All())
	if length != 0 {
		t.Errorf("Unexpected length returned. Expected 0 got %d", length)
	}

	last := slice.Last()
	if last != nil {
		t.Error("Non-nil last value returned for empty ring")
	}
}
