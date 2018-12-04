# Ring

Ring is a threadsafe rotating array in Golang. You define a maximum number of items
and for each added item, it will fill the array up to that max, then begin
to wrap around the start of the array.

## Example

```golang
ring := ring.New(15)

i := 0
for i < 50 {
    ring.Add(i)
    i++
}

// All of the last values (up to the maximum) sorted by when they were added
everything := ring.All()
// The last added item
last := ring.Last()
```
