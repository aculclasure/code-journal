package circular

import (
	"errors"
)

// Buffer represents a circular buffer
type Buffer struct {
	data              []byte
	oldest, next, len int
}

// NewBuffer returns a new Buffer of a specified size.
func NewBuffer(size int) *Buffer {
	if size < 1 {
		return nil
	}
	return &Buffer{data: make([]byte, size)}
}

// ReadByte returns the oldest byte in the given Buffer b or returns
// an error if b is empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.len == 0 {
		return 0, errors.New("cannot read from an empty Buffer")
	}

	data := b.data[b.oldest]
	b.oldest = (b.oldest + 1) % len(b.data)
	b.len--
	return data, nil
}

// WriteByte writes the byte c into the given Buffer b. It returns an error
// if b is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.len == len(b.data) {
		return errors.New("cannot write to full Buffer")
	}

	b.data[b.next] = c
	b.next = (b.next + 1) % len(b.data)
	b.len++
	return nil
}

// Overwrite writes the byte c over the oldest item in the given Buffer b if
// b is full. If b is not full, c is just written into the next available
// position instead.
func (b *Buffer) Overwrite(c byte) {
	if b.len == len(b.data) {
		b.data[b.oldest] = c
		b.oldest = (b.oldest + 1) % len(b.data)
	} else {
		b.WriteByte(c)
	}
}

// Reset resets the given Buffer b to an empty state.
func (b *Buffer) Reset() {
	b.len = 0
}
