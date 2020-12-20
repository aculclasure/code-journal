package react

import (
	"sync"
)

// New returns a new Reactor implemented as an underlying sheet.
func New() Reactor {
	return &sheet{}
}

// sheet represents the collection of different cell types.
type sheet struct {
	inputCells         []*inputCell
	singleComputeCells []*computeCell1
	doubleComputeCells []*computeCell2
}

// CreateInput accepts a number and creates an input cell
// linked into the sheet with the number as its value.
func (s *sheet) CreateInput(v int) InputCell {
	i := &inputCell{value: v, s: s}
	s.inputCells = append(s.inputCells, i)
	return i
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (s *sheet) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	computeCell := &computeCell1{
		value:        f(c.Value()),
		upstreamCell: c,
		s:            s,
		f:            f,
		callbacks:    make(map[callbackID]func(int)),
	}
	s.singleComputeCells = append(s.singleComputeCells, computeCell)
	return computeCell
}

// CreateCompute1 creates a compute cell which computes its value
// based on two other cells. The compute function will only be called
// if the value of either of the passed cells changes.
func (s *sheet) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	computeCell := &computeCell2{
		value:         f(c1.Value(), c2.Value()),
		upstreamCell1: c1,
		upstreamCell2: c2,
		s:             s,
		f:             f,
		callbacks:     make(map[callbackID]func(int)),
	}
	s.doubleComputeCells = append(s.doubleComputeCells, computeCell)
	return computeCell
}

// update updates the computed cells in the sheet.
func (s *sheet) update() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, c := range s.singleComputeCells {
			if c.value != c.Value() {
				c.value = c.Value()
				for _, f := range c.callbacks {
					f(c.value)
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for _, c := range s.doubleComputeCells {
			if c.value != c.Value() {
				c.value = c.Value()
				for _, f := range c.callbacks {
					f(c.value)
				}
			}
		}
	}()

	wg.Wait()
}

// inputCell represents a cell with a changeable value that can
// trigger updates in other cells.
type inputCell struct {
	value int
	s     *sheet
}

// Value returns the value of the input cell.
func (i *inputCell) Value() int {
	return i.value
}

// SetValue accepts a number and sets the value of the inputCell
// to that number.
func (i *inputCell) SetValue(v int) {
	i.value = v
	i.s.update()
}

// callbackID represents a handle to a callback function.
type callbackID *func(int)

// computeCell1 represent a compute cell whose value depends
// on the value of one other cell changing.
type computeCell1 struct {
	value        int
	upstreamCell Cell
	s            *sheet
	f            func(int) int
	callbacks    map[callbackID]func(int)
}

// Value returns the value of the computeCell1.
func (c *computeCell1) Value() int {
	return c.f(c.upstreamCell.Value())
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (c *computeCell1) AddCallback(f func(int)) Canceler {
	c.callbacks[&f] = f
	return &computeCell1Canceler{c: c, callbackID: &f}
}

// computeCell1Canceler provides a Cancel method that allows a
// callback to be deleted from its associated computeCell1
// instance.
type computeCell1Canceler struct {
	c          *computeCell1
	callbackID callbackID
}

// Cancel deletes a callback from its associated computeCell1
// instance.
func (c *computeCell1Canceler) Cancel() {
	delete(c.c.callbacks, c.callbackID)
}

// computeCell2 represents a compute cell whose value depends on
// the value of 2 other cells.
type computeCell2 struct {
	value         int
	upstreamCell1 Cell
	upstreamCell2 Cell
	s             *sheet
	f             func(int, int) int
	callbacks     map[callbackID]func(int)
}

// Value returns the value of the computeCell2.
func (c *computeCell2) Value() int {
	return c.f(c.upstreamCell1.Value(), c.upstreamCell2.Value())
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (c *computeCell2) AddCallback(f func(int)) Canceler {
	c.callbacks[&f] = f
	return &computeCell2Canceler{c: c, callbackID: &f}
}

// computeCell2Canceler provides a Cancel method that allows a
// callback to be deleted from its associated computeCell2
// instance.
type computeCell2Canceler struct {
	c          *computeCell2
	callbackID callbackID
}

// Cancel deletes a callback from its associated computeCell2
// instance.
func (c *computeCell2Canceler) Cancel() {
	delete(c.c.callbacks, c.callbackID)
}
