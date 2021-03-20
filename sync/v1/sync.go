package v1

import "sync"

// Counter will increment a number.
type Counter struct {
	value int
	mu sync.Mutex
}

// Inc the count.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.value
}
