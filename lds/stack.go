package lds

// decided to make it a struct instead of []T, so new fields can easily be added later if needed
// can be rewritten to []T, cause it's such an easy DS to write
type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{[]T{}}
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *Stack[T]) Pop() T {
	l := len(s.arr)
	el := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return el
}

func (s *Stack[T]) Peek() T {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.arr) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}
