package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

// Cache - потокобезопасная структура для хранения данных
type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

// NewCache создает новый экземпляр Cache
func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]int),
	}
}

// Set устанавливает значение по ключу
func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

// Get возвращает значение по ключу
func (c *Cache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

// Increase увеличивает значение по ключу на указанную величину
func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	// Запускаем 10 горутин для увеличения значения
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	// Запускаем 10 горутин для установки значения
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}(i) // Передаем i как параметр
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим конечное значение
	fmt.Println(cache.Get(k1))
}
