package tsp

import (
	"errors"
)

var ErrInsufficientStack = errors.New("insufficient number of elements in the stack")
var ErrOutOfRange = errors.New("out of range")

type Stack struct {
	mem []Value
}

func NewStack() *Stack {
	return &Stack{mem: make([]Value, 0)}
}

func (s *Stack) Push(v Value) {
	s.mem = append(s.mem, v)
}

func (s *Stack) Pop() (Value, error) {
	if len(s.mem) > 0 {
		v := s.mem[len(s.mem) - 1]
		s.mem = s.mem[:len(s.mem) - 1]
		return v, nil
	}
	return 0, ErrInsufficientStack
}

func (s *Stack) Top() (Value, error) {
	if len(s.mem) > 0 {
		return s.mem[len(s.mem) - 1], nil
	}
	return 0, ErrInsufficientStack
}

func (s *Stack) Swap(i int) error {
	l := len(s.mem)
	if 0 < i && i <= l {
		s.mem[l - 1], s.mem[l - i] = s.mem[l - i], s.mem[l - 1]
		return nil
	}
	return ErrOutOfRange
}

func (s *Stack) Dup() error {
	v, err := s.Top()
	if err != nil {
		return err
	}
	s.Push(v)
	return nil
}

func (s *Stack) Sub() error {
	v1, err := s.Pop()
	if err != nil {
		return err
	}
	v2, err := s.Pop()
	if err != nil {
		s.Push(v1)
		return ErrInsufficientStack
	}
	s.Push(v2 - v1)
	return nil
}

func (s *Stack) Greater() error {
	v1, err := s.Pop()
	if err != nil {
		return err
	}
	v2, err := s.Pop()
	if err != nil {
		s.Push(v1)
		return ErrInsufficientStack
	}
	if v2 > v1 {
		s.Push(1)
	} else {
		s.Push(0)
	}
	return nil
}
