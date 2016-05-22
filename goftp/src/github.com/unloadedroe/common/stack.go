package common

type Stack struct {
	top  *Element
	size int64
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Len() int64 {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	var value interface{} = nil
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
	}
	return value
}

func (s *Stack) Peek() interface{} {
	return s.top.value
}

func NewStack() *Stack {
	return new(Stack)
}
