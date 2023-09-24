package lds

const (
	initCapacity = 10
)

// how dumb it is to write a dynamic array in a language, that has slices...?))
type dynamicArray[T any] struct {
	arr []T
	len int
	cap int
}

func NewDArray[T any]() *dynamicArray[T] {
	return &dynamicArray[T]{
		make([]T, initCapacity),
		0,
		10,
	}
}
