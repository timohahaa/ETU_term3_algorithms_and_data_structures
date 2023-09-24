package lds

// decided to make it a struct instead of []T, so new fields can easily be added later if needed
// can be rewritten to []T, cause it's such an easy DS to write
type stack[T any] struct {
	arr []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{[]T{}}
}

func (s *stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *stack[T]) Pop() T {
	l := len(s.arr)
	el := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return el
}

func (s *stack[T]) Peek() T {
	return s.arr[len(s.arr)-1]
}

func (s *stack[T]) Empty() bool {
	return len(s.arr) == 0
}
