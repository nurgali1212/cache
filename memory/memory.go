package memory

import (
	"errors"
	"fmt"
)

type Cache struct {
	memory map[string]interface{}
}

func New() *Cache {
	return &Cache{memory: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) (err error) {
	c.memory[key] = value
	if err != nil {
		fmt.Println("key not found")
	}
	return

}

func (c *Cache) Get(key string) (interface{}, bool) {
	_, ok := c.memory[key]
	if !ok {
		return nil, false
	}
	return c.memory[key], true
}

func (c *Cache) Delete(key string) error {
	_, ok := c.memory[key]
	if !ok {
		err := errors.New("key not found")
		if err != nil {
			return err
		}
	}
	delete(c.memory,key)
	return nil
}
