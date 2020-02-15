package listops

// IntList is a list of integers
type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append returns an IntList that contains the values in newList
// appended to r
func (r IntList) Append(newList IntList) IntList {
	updatedList := make(IntList, r.Length()+newList.Length())
	i := 0

	for _, val := range r {
		updatedList[i] = val
		i++
	}
	for _, val := range newList {
		updatedList[i] = val
		i++
	}
	return updatedList
}

// Concat returns a flattened list that is the concatenation
// of otherLists with r
func (r IntList) Concat(otherLists []IntList) IntList {
	var concatenatedList IntList

	concatenatedList = concatenatedList.Append(r)
	for _, otherList := range otherLists {
		concatenatedList = concatenatedList.Append(otherList)
	}
	return concatenatedList
}

// Length returns the length of the IntList
func (r IntList) Length() int {
	count := 0

	for range r {
		count++
	}
	return count
}

// Reverse returns a list that contains the items of r
// in reversed order
func (r IntList) Reverse() IntList {
	reversedList := make(IntList, r.Length())
	index := 0

	for j := r.Length() - 1; j >= 0; j-- {
		reversedList[index] = r[j]
		index++
	}
	return reversedList
}

// Filter accepts a predicate function argument and returns
// all items in r that match the predicate
func (r IntList) Filter(f predFunc) IntList {
	matchingItems := IntList{}

	for _, v := range r {
		if f(v) {
			matchingItems = matchingItems.Append(IntList{v})
		}
	}
	return matchingItems
}

// Foldl given a list, a function, and an initial accumulator value folds
// each item from the left into the accumulator using the function
func (r IntList) Foldl(f binFunc, accumulator int) int {
	for _, v := range r {
		accumulator = f(accumulator, v)
	}
	return accumulator
}

// Foldr given a list, a function, and an initial accumulator value folds
// each item fromt he right into the accumulator using the function
func (r IntList) Foldr(f binFunc, accumulator int) int {
	for _, v := range r.Reverse() {
		accumulator = f(v, accumulator)
	}
	return accumulator
}

// Map accepts a list of values and a function and returns
// the result of applying the function to each value in the
// list.
func (r IntList) Map(f unaryFunc) IntList {
	results := IntList{}

	for _, v := range r {
		results = results.Append(IntList{f(v)})
	}
	return results
}
