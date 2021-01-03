package binarysearchtree

import "errors"

// Element represents a node in the linked list.
type Element struct {
	data *SearchTreeData
	next *Element
}

// List represents a linked list.
type List struct {
	head *Element
	size int
}

// Size returns the number of items in the List l.
func (l *List) Size() int {
	return l.size
}

// Push appends a new SearchTreeData node s into the List l.
func (l *List) Push(s *SearchTreeData) {
	switch {
	case l.size == 0:
		l.head = &Element{data: s}
	default:
		node := l.head
		for node.next != nil {
			node = node.next
		}
		node.next = &Element{data: s}
	}
	l.size++
}

// Pop removes and returns the last item from the List l or returns
// an error if l is empty.
func (l *List) Pop() (*SearchTreeData, error) {
	switch {
	case l.size == 0:
		return nil, errors.New("cannot Pop from an empty List")
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
