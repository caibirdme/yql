package stack

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

// Init initializes the stack
func (s *BoolStack) Init() {
	if s.sp > -1 {
		s.data = s.data[:0]
	}
	s.sp = -1
}

// NewStack create a new stack object
func NewStack() *BoolStack {
	return &BoolStack{sp: -1}
}
