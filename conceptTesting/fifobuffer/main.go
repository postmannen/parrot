package main

import "fmt"

// ringBuffer holds the data of the buffer,
type ringBuffer struct {
	data []string
}

// newringBuffer is a push/pop storage for values.
func newringBuffer() *ringBuffer {
	return &ringBuffer{}
}

// push will add another item to the end of the buffer with a normal append
func (s *ringBuffer) push(d string) {
	s.data = append(s.data, d)
}

// pop will remove and return the first element of the buffer
func (s *ringBuffer) pop() string {
	if len(s.data) == 0 {
		return ""
	}

	v := s.data[0]
	s.data = append(s.data[0:0], s.data[1:]...)

	return v
}

func main() {
	ts := newringBuffer()
	ts.push("ape")
	ts.push("hest")
	ts.push("ku")

	fmt.Printf("pop = %v\n", ts.pop())
	fmt.Printf("pop = %v\n", ts.pop())

	ts.push("giraff")

	fmt.Printf("pop = %v\n", ts.pop())
	fmt.Printf("pop = %v\n", ts.pop())
	fmt.Printf("pop = %v\n", ts.pop())
	fmt.Printf("pop = %v\n", ts.pop())
}
