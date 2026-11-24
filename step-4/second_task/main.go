package main

import (
	"sync"
)

type Counter struct {
	value int
	mu sync.RWMutex
}

type Count interface {
	Increment()
	Value() int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}
