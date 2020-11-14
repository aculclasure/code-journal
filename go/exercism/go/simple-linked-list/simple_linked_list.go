package linkedlist

import "errors"

// Element represents a node in the linked list.
type Element struct {
	data int
	next *Element
}

// List represents a linked list.
type List struct {
	head *Element
	size int
}

// New accepts a slice of integers and returns a List from them.
func New(data []int) *List {
	list := &List{}

	for _, d := range data {
		list.Push(d)
	}
	return list
}

// Size returns the number of items in the List l.
func (l *List) Size() int {
	return l.size
}

// Push appends a new int n into the List l.
func (l *List) Push(n int) {
	switch {
	case l.size == 0:
		l.head = &Element{data: n}
	default:
		node := l.head
		for node.next != nil {
			node = node.next
		}
		node.next = &Element{data: n}
	}
	l.size++
}

// Pop removes and returns the last item from the List l or returns
// an error if l is empty.
func (l *List) Pop() (int, error) {
	switch {
	case l.size == 0:
		return 0, errors.New("cannot Pop from an empty List")
	case l.size == 1:
		data := l.head.data
		l.head = nil
		l.size--
		return data, nil
	default:
		prev, node := l.head, l.head
		for node.next != nil {
			prev = node
			node = node.next
		}
		data := node.data
		prev.next = nil
		l.size--
		return data, nil
	}
}

// Array returns the items in List l as a slice of integers.
func (l *List) Array() []int {
	a := make([]int, l.size)
	node := l.head

	for i := 0; i < l.size; i++ {
		a[i] = node.data
		node = node.next
	}
	return a
}

// Reverse returns a new List with the items of the given List l in reverse order.
func (l *List) Reverse() *List {
	reversed := &List{}

	for l.size > 0 {
		data, _ := l.Pop()
		reversed.Push(data)
	}
	return reversed
}
