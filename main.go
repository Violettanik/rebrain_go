package main

import (
	"sync"
	"time"
        "fmt"
)

// CacheItem представляет элемент кэша с временем жизни (TTL)
type CacheItem struct {
	Value      interface{}
	Expiration int64 // UnixNano timestamp
}

// Cache - потокобезопасная реализация кэша
type Cache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

// NewCache создает новый экземпляр кэша
func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]CacheItem),
	}
	go c.cleanupExpired() // Запускаем фоновую очистку
	return c
}

// Set добавляет или обновляет значение в кэше
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

// Get возвращает значение по ключу (и флаг его наличия)
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	
	if time.Now().UnixNano() > item.Expiration {
		return nil, false
	}
	
	return item.Value, true
}

// Delete удаляет элемент из кэша
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.items, key)
}

// cleanupExpired периодически удаляет просроченные элементы
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mu.Lock()
		now := time.Now().UnixNano()
		for key, item := range c.items {
			if now > item.Expiration {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

func main() {
	cache := NewCache()
	
	// Запись в кэш
	cache.Set("user:123", "John Doe", 5*time.Minute)
	
	// Чтение из кэша
	if val, ok := cache.Get("user:123"); ok {
		fmt.Println("Found:", val)
	}
	
	// Удаление
	cache.Delete("user:123")
}
