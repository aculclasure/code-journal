package linkedlist

import "errors"

// Node represents a node in the linked list.
type Node struct {
	Val        interface{}
	prev, next *Node
}

// Next returns the next node.
func (e *Node) Next() *Node {
	return e.next
}

// Prev returns the previous node.
func (e *Node) Prev() *Node {
	return e.prev
}

// ErrEmptyList is an error indicating a list operation
// was performed against an empty list.
var ErrEmptyList = errors.New("list cannot be empty")

// List represents a doubly-linked list.
type List struct {
	head, tail *Node
}

// NewList accepts a slice of items and returns a
// doubly-linked list containing the items in the
// same order.
func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{}
	}

	list := &List{}
	for _, v := range args {
		list.PushBack(v)
	}
	return list
}

// PushFront accepts an item v and pushes it to the front
// of the linked list.
func (l *List) PushFront(v interface{}) {
	n := &Node{next: l.head, Val: v}

	if l.head == nil {
		l.tail = n
	} else {
		l.head.prev = n
	}
	l.head = n
}

// PushBack accepts an item v and appends it to the
// linked list.
func (l *List) PushBack(v interface{}) {
	n := &Node{prev: l.tail, Val: v}

	if l.tail == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

// PopFront removes and returns the first item of the linked
// list or returns an error otherwise.
func (l *List) PopFront() (interface{}, error) {
	firstNode := l.First()

	if firstNode == nil {
		return nil, ErrEmptyList
	}

	l.head = firstNode.Next()
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return firstNode.Val, nil
}

// PopBack removes and returns the last item of the linked
// list or returns an error otherwise.
func (l *List) PopBack() (interface{}, error) {
	lastNode := l.Last()

	if lastNode == nil {
		return nil, ErrEmptyList
	}

	l.tail = lastNode.Prev()
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return lastNode.Val, nil
}

// First returns the first Node in the list.
func (l *List) First() *Node {
	return l.head
}

// Last returns the last Node in the list.
func (l *List) Last() *Node {
	return l.tail
}

// Reverse reverses the items in the given list and returns
// the reversed list.
func (l *List) Reverse() {
	for e := l.First(); e != nil; e = e.Prev() {
		e.next, e.prev = e.prev, e.next
	}
	l.head, l.tail = l.tail, l.head
}
