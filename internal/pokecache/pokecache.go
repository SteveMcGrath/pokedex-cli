package pokecache

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
	age     time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		fmt.Println("got from cache")
		return entry.value, true
	}
	return nil, false
}

func (c *Cache) GetJSON(url string, data any) error {
	obj, ok := c.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		obj, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		c.Add(url, obj)
	}
	err := json.Unmarshal(obj, data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) reapLoop() {
	go func() {
		for {
			time.Sleep(c.age)
			now := time.Now()
			c.mu.Lock()
			for k, entry := range c.entries {
				age := entry.createdAt.Add(c.age)

				if now.After(age) || now.Equal(age) {
					delete(c.entries, k)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func NewCache() *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		age:     time.Second * 5,
	}
	c.reapLoop()
	return &c
}
