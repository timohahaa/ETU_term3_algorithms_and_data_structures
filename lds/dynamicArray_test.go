package lds

import (
	"testing"
)

func TestGetDA(t *testing.T) {
	a := NewDArray[int]()
	_, err := a.Get(228)
	if err == nil {
		t.Errorf("should have gotten an error - a.Get(228) - the array is empty")
	}
	a.Add(0, 8)
	_, err = a.Get(0)
	if err != nil {
		t.Errorf("should have no error - a.Get(0)")
	}
}

func TestDeleteDA(t *testing.T) {
	a := NewDArray[int]()
	a.PushBack(8)
	a.PushBack(1337)
	a.PushBack(10)
	_, err := a.Get(1)
	if err != nil {
		t.Errorf("should have gotten no error - index exists")
	}
	a.Delete(1)
	val, err := a.Get(1)
	if err != nil {
		t.Errorf("should have gotten no error - a.Get(1) - index exists")
	}
	if val == 1337 {
		t.Errorf("value was not deleted")
	}
}

func binarySearchLowerBound[T any](arr DynamicArray[T], cmp func(a, b T) int, target T) int {
	low, high := 0, arr.Len()
	for low < high {
		mid := (low + high) / 2
		if cmp(arr.GetNoError(mid), target) >= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low < arr.Len() { // && cmp(arr[low], target) == 0 {
		return low
	}
	return -1
}

func TestLen(t *testing.T) {
	a := NewDArray[int]()
	if a.Len() != 0 {
		t.Errorf("wrong length: %v, need 0", a.Len())
	}
	a.PushBack(1)
	if a.Len() != 1 {
		t.Errorf("wrong length: %v, need 1", a.Len())
	}
	a.PushBack(2)
	if a.Len() != 2 {
		t.Errorf("wrong length: %v, need 2", a.Len())
	}

	if a.GetNoError(0) != 1 {
		t.Errorf("wrong result: %v, need 0", a.GetNoError(0))
	}
	if a.GetNoError(1) != 2 {
		t.Errorf("wrong result: %v, need 2", a.GetNoError(1))
	}
}
