package paasio

import (
	"io"
	"sync"
)

type counter struct {
	nbytes int64
	nops   int
	mutex  sync.Mutex
}

func (c *counter) update(nbytes int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.nbytes += nbytes
	c.nops++
}

func (c *counter) getCount() (int64, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.nbytes, c.nops
}

type readCounter struct {
	reader  io.Reader
	counter *counter
}

// NewReadCounter returns a ReadCounter that wraps the given io.Reader r.
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{reader: r, counter: &counter{}}
}

// Read reads a slice of bytes and returns the number of bytes read or an error.
func (r *readCounter) Read(p []byte) (int, error) {
	nbytes, err := r.reader.Read(p)

	if err != nil {
		return 0, err
	}
	r.counter.update(int64(nbytes))
	return nbytes, nil
}

// ReadCount returns the number of bytes read and the number of reads.
func (r *readCounter) ReadCount() (int64, int) {
	return r.counter.getCount()
}

type writeCounter struct {
	writer  io.Writer
	counter *counter
}

// NewWriteCounter returns a WriteCounter on the given io.Writer w.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{writer: w, counter: &counter{}}
}

// Write writes a slice of bytes and returns the number of bytes written or an error.
func (w *writeCounter) Write(p []byte) (int, error) {
	nbytes, err := w.writer.Write(p)

	if err != nil {
		return 0, err
	}
	w.counter.update(int64(nbytes))
	return nbytes, nil
}

// WriteCount returns the number of bytes written and the number of writes.
func (w *writeCounter) WriteCount() (int64, int) {
	return w.counter.getCount()
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

// NewReadWriteCounter returns a ReadWriteCounter on the given io.ReadWriter rw.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter{reader: rw, counter: &counter{}},
		writeCounter{writer: rw, counter: &counter{}},
	}
}
