package stack

import (
	"sync"
)

// BoolStack implement a basic stack
type BoolStack struct {
	data []bool
	sp   int
}

// Pop pops the element atop of the stack
func (s *BoolStack) Pop() bool {
	if s.sp < 0 {
		return false
	}
	r := s.data[s.sp]
	s.data = s.data[:s.sp]
	s.sp--
	return r
}

// Push pushes an element into the top of the stack
func (s *BoolStack) Push(b bool) {
	s.sp++
	s.data = append(s.data, b)
}

// NewStack create a new stack object
// Make sure invoking release to avoid memory leak
func NewStack() (stack *BoolStack, release func()) {
	b := pool.Get().(*BoolStack)
	if b.sp > -1 {
		b.sp = -1
		b.data = b.data[:0]
	}
	return b, func() { pool.Put(b) }
}

var pool = sync.Pool{
	New: func() interface{} {
		return &BoolStack{sp: -1}
	},
}
