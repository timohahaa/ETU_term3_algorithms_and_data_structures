package lds

import "testing"

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
