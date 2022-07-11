package cache

import "time"

type Entry struct {
	value    string
	deadline time.Time
}

type Cache struct {
	entries map[string]Entry
}

func NewCache() Cache {
	return Cache{entries: make(map[string]Entry)}
}

func (c *Cache) Get(key string) (string, bool) {
	c.deleteExpiredKeys()

	entry, exists := c.entries[key]

	if !exists {
		return "", false
	}
	return entry.value, true

}

func (c *Cache) Put(key, value string) {
	c.deleteExpiredKeys()

	c.entries[key] = Entry{value: value}
}

func (c *Cache) Keys() []string {
	c.deleteExpiredKeys()

	keys := make([]string, 0, len(c.entries))

	for k := range c.entries {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.deleteExpiredKeys()

	c.entries[key] = Entry{value: value, deadline: deadline}
}

func (c *Cache) deleteExpiredKeys() {
	now := time.Now()

	for key, entry := range c.entries {
		if !entry.deadline.IsZero() && entry.deadline.Before(now) {
			delete(c.entries, key)
		}
	}
}
