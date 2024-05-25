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
	Add(key string, value interface{}) error
	Delete(key string) error
}

type inMemoryCache struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}

type cacheItem struct {
	value interface{}
}

func NewInMemoryCache() Cache {
	return &inMemoryCache{
		data: make(map[string]cacheItem),
	}
}

func (c *inMemoryCache) FillUp(data map[string]models.Order) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = make(map[string]cacheItem)
	for key, value := range data {
		c.data[key] = cacheItem{value: value}
	}
	return nil
}

func (c *inMemoryCache) Set(key string, value interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheItem{
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

func (c *inMemoryCache) Add(key string, value interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, found := c.data[key]; found {
		return errors.New("item already exists")
	}
	c.data[key] = cacheItem{
		value: value,
	}
	return nil
}

func (c *inMemoryCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, found := c.data[key]; !found {
		return errors.New("item not found")
	}
	delete(c.data, key)
	return nil
}
