package lds

import (
	"errors"
	"fmt"
)

const (
	initCapacity = 10
)

// how dumb it is to write a dynamic array in a language, that has slices...?))
type DynamicArray[T any] struct {
	arr []T
	len int
	cap int
}

func NewDArray[T any]() *DynamicArray[T] {
	return &DynamicArray[T]{
		arr: make([]T, initCapacity),
		len: 0,
		cap: 10,
	}
}

func (a *DynamicArray[T]) Len() int {
	return a.len
}

func (a *DynamicArray[T]) Cap() int {
	return a.cap
}

func (a *DynamicArray[T]) EnsureCapacity(newCap int) {
	if a.cap >= newCap {
		return
	}
	// always grow by the factor of two
	newArr := make([]T, newCap*2)
	// can also do it by hand
	copy(newArr, a.arr)
	a.arr = newArr
	a.cap = newCap
}

func (a *DynamicArray[T]) Get(idx int) (T, error) {
	var zeroVal T
	if idx >= a.len || idx < 0 {
		return zeroVal, errors.New("index out of range")
	}
	return a.arr[idx], nil
}

// get method, that returns no error
// panics at an invalid index
func (a *DynamicArray[T]) GetNoError(idx int) T {
	return a.arr[idx]
}

func (a *DynamicArray[T]) Set(idx int, data T) {
	if idx >= a.len || idx < 0 {
		return
	}
	a.arr[idx] = data
}

// if index is out of range, this function does nothing
func (a *DynamicArray[T]) Add(idx int, data T) {
	if idx > a.len || idx < 0 {
		return
	}
	a.EnsureCapacity(a.len + 1)
	// push front
	if idx == 0 {
		copy(a.arr, a.arr[1:])
		a.arr[0] = data
		a.len++
		return
	}
	// push back
	if idx == a.len {
		a.arr[a.len] = data
		a.len++
		return
	}
	// add in general
	copy(a.arr[idx+1:], a.arr[idx:])
	a.arr[idx] = data
	a.len++
}

// if index is out of range, this function does nothing
func (a *DynamicArray[T]) Delete(idx int) {
	if idx >= a.len || idx < 0 {
		return
	}
	copy(a.arr[idx:], a.arr[idx+1:])
	a.len--
}

func (a *DynamicArray[T]) PushBack(data T) {
	a.Add(a.len, data)
}

func (a *DynamicArray[T]) Print() {
	for i, val := range a.arr {
		fmt.Printf("[%v]: %+v\n", i, val)
	}
}
