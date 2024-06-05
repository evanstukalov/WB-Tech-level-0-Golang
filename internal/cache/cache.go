package cache

import (
	"errors"
	"github.com/evanstukalov/wildberries_internship_l0/internal/models"
	"sync"
)

type Cache interface {
	FillUp(data map[string]models.Order) error
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type inMemoryCache struct {
	data  map[string]item
	mutex sync.RWMutex
}

type item struct {
	value interface{}
}

func NewInMemoryCache() Cache {
	return &inMemoryCache{
		data: make(map[string]item),
	}
}

func (c *inMemoryCache) FillUp(data map[string]models.Order) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]item)
	for key, value := range data {
		c.data[key] = item{value: value}
	}
	return nil
}

func (c *inMemoryCache) Set(key string, value interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = item{
		value: value,
	}
	return nil
}

func (c *inMemoryCache) Get(key string) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.data[key]
	if !found {
		return nil, errors.New("item not found")
	}

	return item.value, nil
}
