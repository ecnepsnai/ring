package ring_test

import (
	"fmt"

	"github.com/ecnepsnai/ring"
)

func ExampleNew() {
	// Make a new ring
	list := ring.New(5)

	// Add something to it
	list.Add("foo")

	// Get the last object added
	fmt.Printf("%s", list.Last())
	// output: foo
}

func ExampleRing_Add() {
	// Make a new ring
	list := ring.New(5)

	// Add something to it
	list.Add("foo")

	// Get the last object added
	fmt.Printf("%s", list.Last())
	// output: foo
}

func ExampleRing_Last() {
	// Make a new ring
	list := ring.New(5)

	// Add something to it
	list.Add("foo")

	// Get the last object added
	fmt.Printf("%s", list.Last())
	// output: foo
}

func ExampleRing_All() {
	// Make a new ring with a maximum capacity of 5 objects
	list := ring.New(5)

	// Add 10 objects (0 through 9)
	i := 0
	for i < 10 {
		list.Add(i)
		i++
	}

	// Get all of the objects form the ring, which will be the last 5 numbers
	// we added
	last5 := list.All()
	fmt.Printf("%v", last5)
	// output: [5 6 7 8 9]
}

func ExampleRing_Truncate() {
	// Make a new ring
	list := ring.New(5)

	// Add 10 objects (0 through 9)
	i := 0
	for i < 10 {
		list.Add(i)
		i++
	}

	// Get the number of objects in the ring
	beforeCount := len(list.All())

	// Truncate the ring
	list.Truncate()

	// Get the number of objects in the ring
	afterCount := len(list.All())

	fmt.Printf("Before: %d, After: %d", beforeCount, afterCount)
	// output: Before: 5, After: 0
}
