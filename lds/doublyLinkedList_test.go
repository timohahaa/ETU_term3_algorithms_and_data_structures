package lds

import "testing"

func TestGetDLL(t *testing.T) {
	l := NewDLList[int]()
	_, err := l.Get(228)
	if err == nil {
		t.Errorf("should have gotten an error - l.Get(0) - the list is empty")
	}
	l.Add(0, 8)
	_, err = l.Get(0)
	if err != nil {
		t.Errorf("should have gotten no error - l.Get(0)")
	}
}

func TestDeleteDLL(t *testing.T) {
	l := NewDLList[int]()
	l.AddTail(8)
	l.AddTail(9)
	l.Add(1, 1337)
	_, err := l.Get(1)
	if err != nil {
		t.Errorf("should have gotten no error - l.Get(1) - index exists")
	}
	// now delete node
	l.Delete(1)
	n, err := l.Get(1)
	if err != nil {
		t.Errorf("should have gotten no error - l.Get(1) - index exists")
	}
	if n.Val == 1337 {
		t.Errorf("value was not deleted")
	}
}
