package tree

import (
	"fmt"
	"sort"
)

// Record represents a record retrieved from some datastore.
type Record struct {
	ID     int
	Parent int
}

// Node is the base unit that is used to build a tree. It contains
// an ID and a list of child Nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Build accepts a slice of Records and returns the root Node
// of a tree built from the Records or an error if any of
// the records is invalid.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	tree := map[int]*Node{}

	for i, record := range records {
		if record.Parent > record.ID {
			return nil, fmt.Errorf(
				"record's id (%d) should be greater than or equal to its parent's id (%d)",
				record.ID, record.Parent)
		}
		if record.ID != i {
			return nil, fmt.Errorf("record should have an id of %d, got %d instead", i, record.ID)
		}
		if record.ID > 0 && record.Parent == record.ID {
			return nil, fmt.Errorf("non-root record with id %d should not refer to itself as a parent",
				record.ID)
		}

		node := &Node{ID: record.ID}
		tree[record.ID] = node
		if record.ID > 0 {
			parent := tree[record.Parent]
			parent.Children = append(parent.Children, node)
		}
	}
	return tree[0], nil
}
