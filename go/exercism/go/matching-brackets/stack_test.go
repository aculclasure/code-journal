// +build stack

package brackets

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	got := newStack()
	want := &stack{items: []rune{}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("newStack() gave %v, want %v", got, want)
	}
}

var pushTestCases = []struct {
	description   string
	initialStack  *stack
	itemsToPush   []rune
	expectedStack *stack
}{
	{
		description:   "Push items into empty stack",
		initialStack:  newStack(),
		itemsToPush:   []rune{'a', 'b', 'c'},
		expectedStack: &stack{items: []rune{'a', 'b', 'c'}},
	},
	{
		description:   "Push items into non-empty stack",
		initialStack:  &stack{items: []rune{'a', 'b', 'c'}},
		itemsToPush:   []rune{'d', 'e', 'f'},
		expectedStack: &stack{items: []rune{'a', 'b', 'c', 'd', 'e', 'f'}},
	},
}

func TestPush(t *testing.T) {
	for _, testCase := range pushTestCases {
		for _, itemToPush := range testCase.itemsToPush {
			testCase.initialStack.push(itemToPush)
		}
		if !reflect.DeepEqual(testCase.initialStack, testCase.expectedStack) {
			t.Fatalf("After pushing %q, expected stack to be %v but got %v instead",
				testCase.itemsToPush, testCase.expectedStack, testCase.initialStack)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

var popTestCases = []struct {
	description   string
	initialStack  *stack
	numItemsToPop int
	expectError   bool
	expectedStack *stack
}{
	{
		description:   "Pop single item from non-empty stack",
		initialStack:  &stack{items: []rune{'a'}},
		numItemsToPop: 1,
		expectError:   false,
		expectedStack: &stack{items: []rune{}},
	},
	{
		description:   "Pop single item from empty stack",
		initialStack:  newStack(),
		numItemsToPop: 1,
		expectError:   true,
		expectedStack: nil,
	},
	{
		description:   "Pop multiple items from non-empty stack",
		initialStack:  &stack{items: []rune{'a', 'b', 'c', 'd', 'e', 'f'}},
		numItemsToPop: 3,
		expectError:   false,
		expectedStack: &stack{items: []rune{'a', 'b', 'c'}},
	},
	{
		description:   "Pop multiple items from empty stack",
		initialStack:  newStack(),
		numItemsToPop: 2,
		expectError:   true,
		expectedStack: nil,
	},
	{
		description:   "Pop more items than available in stack",
		initialStack:  &stack{items: []rune{'a', 'b', 'c'}},
		numItemsToPop: 4,
		expectError:   true,
		expectedStack: nil,
	},
}

func TestPop(t *testing.T) {
	for _, testCase := range popTestCases {
		var err error
		var lastPoppedItem rune
		poppedItems := []rune{}

		for i := 0; i < testCase.numItemsToPop; i++ {
			lastPoppedItem, err = testCase.initialStack.pop()
			if err != nil {
				poppedItems = append(poppedItems, lastPoppedItem)
			}
		}
		if testCase.expectError {
			if err == nil {
				t.Fatalf("calling pop() %d times should have returned an error, got %v instead",
					testCase.numItemsToPop, poppedItems)
			}
		} else {
			if err != nil {
				t.Fatalf("calling pop() %d times produced an unexpected error: %v",
					testCase.numItemsToPop, err.Error())
			} else {
				if !reflect.DeepEqual(testCase.initialStack, testCase.expectedStack) {
					t.Fatalf("After popping %d items, got stack %v, expected %v",
						testCase.numItemsToPop, testCase.initialStack, testCase.expectedStack)
				}
			}
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

var isEmptyTestCases = []struct {
	description  string
	initialStack *stack
	expected     bool
}{
	{
		description:  "Empty stack",
		initialStack: newStack(),
		expected:     true,
	},
	{
		description:  "Non-empty stack",
		initialStack: &stack{items: []rune{'a', 'b', 'c'}},
		expected:     false,
	},
}

func TestIsEmpty(t *testing.T) {
	for _, testCase := range isEmptyTestCases {
		got := testCase.initialStack.isEmpty()
		if got != testCase.expected {
			t.Fatalf("isEmpty() for stack %v returned %v, want %v",
				testCase.initialStack, got, testCase.expected)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}
