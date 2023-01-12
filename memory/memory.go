package memory

import (
	"errors"
	"fmt"
	"time"
)

type Cache struct {
	memory map[string]itemCache
	ttl    time.Duration
}
type itemCache struct {
	value          interface{}
	timeDeleteUnix int64
}

func New() *Cache {
	return &Cache{memory: make(map[string]itemCache)}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.memory[key] = itemCache{
		value:          value,
		timeDeleteUnix: time.Now().Add(duration).Unix(),
	}


}

func (c *Cache) Get(key string) (interface{}, error) {
	itemCache, ok := c.memory[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %s not found", key))
	}
	if itemCache.timeDeleteUnix < time.Now().Unix() {
		return nil, errors.New(fmt.Sprintf("key %s is outdated", key))
	}
	return itemCache.value, nil
}

func (c *Cache) Delete(key string) error {
	_, ok := c.memory[key]
	if !ok {
		err := errors.New("key not found")
		if err != nil {
			return err
		}
	}
	delete(c.memory, key)
	return nil
}
