package lds // list data structures

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	Val  T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	len  int
	head *node[T]
	tail *node[T]
}

func NewDLList[T any]() *list[T] {
	return &list[T]{0, nil, nil}
}

func (l *list[T]) Len() int {
	return l.len
}

func (l *list[T]) Head() *node[T] {
	return l.head
}

func (l *list[T]) Tail() *node[T] {
	return l.tail
}

func (l *list[T]) Get(idx int) (*node[T], error) {
	cur := l.head
	if idx >= l.len || idx < 0 {
		return nil, errors.New("index out of range")
	}
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	return cur, nil
}

// if index is out of range, this function does nothing
func (l *list[T]) Add(idx int, data T) {
	node := &node[T]{data, nil, nil}
	if idx > l.len || idx < 0 {
		// do nothing
		return
	}
	// check if list is not empty
	if l.head == nil {
		l.head = node
		l.tail = node
		l.len++
		return
	}
	// add at head
	if idx == 0 {
		node.next = l.head
		l.head.prev = node
		l.head = node
		l.len++
		return
	}
	// add at tail
	if idx == l.len {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
		l.len++
		return
	}
	// add in general
	prev, _ := l.Get(idx - 1)
	node.next = prev.next
	prev.next.prev = node
	prev.next = node
	node.prev = prev
	l.len++
}

// if index is out of range, this function does nothing
func (l *list[T]) Delete(idx int) {
	if idx >= l.len || idx < 0 {
		// do nothing
		return
	}
	// delete at head
	if idx == 0 {
		// old head should be garbage-collected I guess
		l.head = l.head.next
		l.head.prev = nil
		if l.head == nil {
			l.tail = nil
		}
		l.len--
		return
	}
	// delete at tail
	if idx == l.len-1 {
		// old tail is garbage-collected
		l.tail = l.tail.prev
		l.len--
		return
	}
	// deleted node should be garbage-collected also
	prev, _ := l.Get(idx - 1)
	nextNext := prev.next.next
	prev.next = nextNext
	nextNext.prev = prev
	l.len--
}

func (l *list[T]) AddHead(data T) {
	l.Add(0, data)
}

func (l *list[T]) AddTail(data T) {
	l.Add(l.len, data)
}

func (l *list[T]) DeleteHead(data T) {
	l.Delete(0)
}

func (l *list[T]) DeleteTail(data T) {
	l.Delete(l.len - 1)
}

func (l *list[T]) Print() {
	cur := l.head
	for cur != nil {
		fmt.Printf("%+v -> ", cur.Val)
		cur = cur.next
	}
}
