package binarysearchtree

import (
	"fmt"
	"log"
)

// SearchTreeData represents a node of a binary search tree.
type SearchTreeData struct {
	left, right *SearchTreeData
	data        int
}

// Bst accepts a number and returns a new binary search tree
// with that number as the root node element.
func Bst(v int) *SearchTreeData {
	return &SearchTreeData{data: v}
}

// Insert accepts a number and inserts it in the appropriate
// location of the binary search tree.
func (s *SearchTreeData) Insert(v int) {
	for next := s; next != nil; {
		if v <= next.data && next.left == nil {
			next.left = &SearchTreeData{data: v}
			break
		}

		if v <= next.data && next.left != nil {
			next = next.left
		}

		if v > next.data && next.right == nil {
			next.right = &SearchTreeData{data: v}
			break
		}

		if v > next.data && next.right != nil {
			next = next.right
		}
	}
}

// MapString accepts a function f and applies f to each
// item in the binary search tree in order and returns the
// result as a list of strings.
func (s *SearchTreeData) MapString(f func(int) string) []string {
	orderedItems, err := s.inOrder()
	if err != nil {
		log.Fatal(err)
	}

	mappedItems := make([]string, len(orderedItems))
	for i, v := range orderedItems {
		mappedItems[i] = f(v.data)
	}
	return mappedItems
}

// MapInt accepts a function f and applies f to each
// item in the binary search tree in order and returns
// the result as a list of integers.
func (s *SearchTreeData) MapInt(f func(int) int) []int {
	orderedItems, err := s.inOrder()
	if err != nil {
		log.Fatal(err)
	}

	mappedItems := make([]int, len(orderedItems))
	for i, v := range orderedItems {
		mappedItems[i] = f(v.data)
	}
	return mappedItems
}

// inOrder returns a list of binary search tree nodes in order
// or returns an error if a problem happens when arranging
// the items in order.
func (s *SearchTreeData) inOrder() ([]*SearchTreeData, error) {
	var (
		items   []*SearchTreeData
		stack   = &List{}
		current = s
		err     error
	)

	for current != nil || stack.Size() > 0 {
		for current != nil {
			stack.Push(current)
			current = current.left
		}
		current, err = stack.Pop()
		if err != nil {
			return nil, fmt.Errorf("got error trying to pop the stack: %v", err)
		}
		items = append(items, current)
		current = current.right
	}
	return items, nil
}
