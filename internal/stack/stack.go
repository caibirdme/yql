package stack

type Stack interface {
	Push(bool)
	Pop() bool
	Init()
}

type baseStack struct {
	data []bool
	sp   int
}

func (s *baseStack) Pop() bool {
	if s.sp < 0 {
		return false
	}
	r := s.data[s.sp]
	s.sp--
	return r
}

func (s *baseStack) Push(b bool) {
	s.sp++
	length := len(s.data)
	if length-1 >= s.sp {
		s.data[s.sp] = b
	} else {
		s.data = append(s.data, b)
	}
}

func (s *baseStack) Init() {
	if s.sp > -1 {
		s.data = s.data[:0]
	}
	s.sp = -1
}

func NewStack() Stack {
	return &baseStack{sp: -1}
}
