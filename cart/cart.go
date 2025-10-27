package cart

import (
	"sync"
)

type Cart struct {
	items map[string]float64 // map of item name to price (could also be quantity: for simplicity, just price per item)
	mu    sync.RWMutex
}

func NewCart() *Cart {
	return &Cart{
		items: make(map[string]float64),
	}
}

// AddItem adds an item to the cart. If already present, replace the price (or sum as needed).
func (c *Cart) AddItem(item string, price float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[item] = price
}

// RemoveItem removes an item from the cart. It is safe if the item does not exist.
func (c *Cart) RemoveItem(item string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, item)
}

// Total returns the sum of all item prices in the cart.
func (c *Cart) Total() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	total := 0.0
	for _, price := range c.items {
		total += price
	}
	return total
}

// Count returns the number of unique items in the cart.
func (c *Cart) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// Snapshot returns a copy of current items for safe read access (for debug or displaying cart state)
func (c *Cart) Snapshot() map[string]float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	copyMap := make(map[string]float64, len(c.items))
	for k, v := range c.items {
		copyMap[k] = v
	}
	return copyMap
}
