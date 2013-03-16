package stack

import (
	"errors"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func New() *Stack {
	return &Stack{top: nil, size: 0}
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Empty() bool {
	if s.size != 0 {
		return false
	}
	return true
}

func (s *Stack) Push(value interface{}) error {
	s.top = &Element{value, s.top}
	s.size++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.size > 0 {
		val := s.top.value
		s.top = s.top.next
		s.size--
		return val, nil
	}
	return 0, errors.New("Unable to Pop: Empty Stack")
}
