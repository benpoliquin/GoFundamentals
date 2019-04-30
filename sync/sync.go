package main

import "sync"

// Counter contains a mutex lock and value int
type Counter struct {
	value int
	mu    sync.Mutex
}

// Inc increments value by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the value of the counter int
func (c *Counter) Value() int {
	return c.value
}

// NewCounter returns a new Counter
func NewCounter() *Counter {
	return &Counter{}
}
