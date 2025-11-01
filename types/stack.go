package types

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidStackOperation = errors.New("invalid stack operation: stack is empty")

type Stack[T any] struct {
	top  int
	cap  int
	data []*T
}

func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{
		top:  -1,
		cap:  cap,
		data: make([]*T, 0, cap),
	}
}

func (s *Stack[T]) Len() int {
	return s.top + 1
}

func (s *Stack[T]) Top() (*T, error) {
	if s.top == -1 {

		return nil, ErrInvalidStackOperation
	}
	return s.data[s.top], nil
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, &v)
	s.top++
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.top == -1 {
		return nil, ErrInvalidStackOperation
	}

	v := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return v, nil
}

func (s *Stack[T]) Clear() {
	s.top = 0
	s.cap = 0
	s.data = make([]*T, 0)
}

func (s Stack[T]) String() string {
	b := strings.Builder{}
	b.WriteString("Stack[")
	for _, v := range s.data {
		b.WriteString(fmt.Sprintf("%v, ", v))
	}
	b.WriteString("]")

	return b.String()
}
