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
